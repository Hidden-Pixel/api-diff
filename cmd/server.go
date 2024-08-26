package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run HTTP and gRPC service",
	Long:  `Run HTTP and gRPC service`,
	Run:   RunServer,
}

type WrappedWriter struct {
	http.ResponseWriter
	StatusCode int
}

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &WrappedWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func CreateHTTPServer() *HTTPServer {
	s := HTTPServer{
		Router: http.NewServeMux(),
	}
	return &s
}

type HTTPServer struct {
	Router *http.ServeMux
}

func (s *HTTPServer) AttachRoutes() {
	v1 := http.NewServeMux()
	v1.Handle("/api", http.StripPrefix("/v1", s.Router))
	s.Router.HandleFunc("GET /version", s.GETVersionHandler)
	s.Router.HandleFunc("GET /request", s.GETRequestHandler)
	s.Router.HandleFunc("GET /diff", s.GETDiffHandler)
}

func (s *HTTPServer) GETVersionHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"version": "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) GETRequestHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"request": "Handled",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) GETDiffHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"diff": "No changes",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) RunServer() {
	middleware := CreateStack(
		Logging,
	)
	server := &http.Server{
		Addr:    ":8081",
		Handler: middleware(s.Router),
	}

	// Channel to listen for interrupt or terminate signal from the OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Println("server listening on port 8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on port 8081: %v\n", err)
		}
	}()

	// Block until a signal is received.
	<-quit
	log.Println("shutting down server...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exiting")
}

func RunServer(cmd *cobra.Command, args []string) {
	server := CreateHTTPServer()
	server.AttachRoutes()
	server.RunServer()
}
