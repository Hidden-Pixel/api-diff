package cmd

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

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
	Configs []Config
}

type Config struct {
	Path   string `mapstructure:"path"`
	Target string `mapstructure:"target"`
}

func CreateReverseProxy() *ReverseProxy {
	var configs []Config
	if err := viper.UnmarshalKey("server", &configs); err != nil {
		log.Fatalf("Unable to decode server configurations: %v", err)
	}
	rp := ReverseProxy{
		Configs: configs,
	}
	for _, s := range configs {
		log.Printf("Setting up proxy for path %s to target %s\n", s.Path, s.Target)
		targetURL, err := url.Parse(s.Target)
		if err != nil {
			log.Fatalf("Failed to parse target URL %s: %v", s.Target, err)
		}
		http.HandleFunc(s.Path, CreateProxyHandler(targetURL))
	}
	return &rp
}

func CreateProxyHandler(target *url.URL) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO(nick): metrics and payloads should be captured here?
		r.URL.Path = "" // reset path so the pull path is proxied
		proxy.ServeHTTP(w, r)
	}
}

func RunReverseProxyServer(cmd *cobra.Command, args []string) {
	CreateReverseProxy()
	port := ":8080"
	log.Printf("Starting reverse proxy server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
