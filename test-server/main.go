package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We need two wait groups for two servers

	// Server 1
	go func() {
		defer wg.Done()
		http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("got request on 8081")
			fmt.Fprintf(w, "Response from server on port 8081")
		})

		fmt.Println("Starting server on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Server 1 failed: %v", err)
		}
	}()

	// Server 2
	go func() {
		defer wg.Done()

		// We need a new ServeMux for the second server to avoid conflict with the default one
		mux := http.NewServeMux()
		mux.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("got request on 8082")
			fmt.Fprintf(w, "Response from server on port 8082")
		})

		fmt.Println("Starting server on :8082")
		if err := http.ListenAndServe(":8082", mux); err != nil {
			log.Fatalf("Server 2 failed: %v", err)
		}
	}()

	// Wait for both servers to run
	wg.Wait()
}
