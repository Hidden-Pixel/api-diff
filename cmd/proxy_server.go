package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reverseProxyServerCmd = &cobra.Command{
	Use:   "reverse-proxy-serve",
	Short: "Run Reverse Proxy",
	Long:  `Run Reverse Proxy`,
	Run:   RunReverseProxyServer,
}

type ReverseProxy struct {
	servers []ServerConfig
}

type ServerConfig struct {
	Path    string         `mapstructure:"path"`
	Targets []TargetConfig `mapstructure:"target"`
}

type TargetConfig struct {
	URL        string `mapstructure:"url"`
	Returnable bool   `mapstructure:"returnable"`
}

func CreateReverseProxy() *ReverseProxy {
	var servers []ServerConfig
	if err := viper.UnmarshalKey("server", &servers); err != nil {
		log.Fatalf("Unable to decode server configurations: %v", err)
	}

	for _, server := range servers {
		fmt.Printf("Setting up proxy for path %s to targets %v\n", server.Path, server.Targets)

		var targetConfigs []TargetConfig
		for _, target := range server.Targets {
			targetConfigs = append(targetConfigs, target)
		}

		http.HandleFunc(server.Path, func(targetConfigs []TargetConfig) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				var wg sync.WaitGroup

				// Channel for collecting the response from the returnable target
				returnableResponseChan := make(chan *http.Response, 1)
				defer close(returnableResponseChan)

				for _, targetConfig := range targetConfigs {
					targetURL, err := url.Parse(targetConfig.URL)
					if err != nil {
						log.Printf("Failed to parse target URL %s: %v", targetConfig.URL, err)
						continue
					}

					wg.Add(1)
					go func(target *url.URL, isReturnable bool) {
						defer wg.Done()

						// Create a copy of the original request
						proxyReq, err := http.NewRequest(r.Method, target.String()+r.URL.Path, r.Body)
						if err != nil {
							log.Printf("Failed to create request for target %s: %v", target, err)
							return
						}
						proxyReq.Header = r.Header

						client := &http.Client{}

						if isReturnable {
							// Send request to the returnable target and send response to the channel
							resp, err := client.Do(proxyReq)
							if err != nil {
								log.Printf("Error forwarding request to returnable target %s: %v", target, err)
								return
							}
							returnableResponseChan <- resp
						} else {
							// Set a context with a 1-minute deadline for non-returnable requests
							ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
							defer cancel()

							// Attach the context to the request
							proxyReq = proxyReq.WithContext(ctx)

							// Send the request (we do not need to collect the response)
							resp, err := client.Do(proxyReq)
							if err != nil {
								log.Printf("Error forwarding request to target %s: %v", target, err)
								return
							}
							defer resp.Body.Close()
						}
					}(targetURL, targetConfig.Returnable)
				}

				// Use the response from the returnable target as soon as it's available
				select {
				case returnableResponse := <-returnableResponseChan:
					if returnableResponse != nil {
						defer returnableResponse.Body.Close()
						body, err := ioutil.ReadAll(returnableResponse.Body)
						if err != nil {
							http.Error(w, "Failed to read response body", http.StatusInternalServerError)
							return
						}
						for k, v := range returnableResponse.Header {
							for _, value := range v {
								w.Header().Add(k, value)
							}
						}
						w.WriteHeader(returnableResponse.StatusCode)
						w.Write(body)
					} else {
						http.Error(w, "No response from returnable target", http.StatusBadGateway)
					}
				case <-time.After(30 * time.Second): // Optional timeout for returnable response
					http.Error(w, "Returnable target did not respond in time", http.StatusGatewayTimeout)
				}
			}
		}(targetConfigs))
	}

	return &ReverseProxy{servers: servers}
}

func RunReverseProxyServer(cmd *cobra.Command, args []string) {
	CreateReverseProxy()
	port := ":8080"
	fmt.Printf("Starting reverse proxy server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
