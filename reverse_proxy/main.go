package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Backend server URLs to forward requests
	targetURL1 := "https://httpbin.org"
	targetURL2 := "https://postman-echo.com"

	// Parse backend URLs
	backendURL1, err := url.Parse(targetURL1)
	if err != nil {
		log.Fatalf("Error parsing backend URL 1: %v", err)
	}
	backendURL2, err := url.Parse(targetURL2)
	if err != nil {
		log.Fatalf("Error parsing backend URL 2: %v", err)
	}

	// Create reverse proxies
	proxy1 := httputil.NewSingleHostReverseProxy(backendURL1)
	proxy2 := httputil.NewSingleHostReverseProxy(backendURL2)

	// Handler for "/some-path1"
	http.HandleFunc("/somepath1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request for:", r.URL.Path)
		r.URL.Path = "/" // Modify the path before forwarding (optional)
		proxy1.ServeHTTP(w, r)
	})

	// Handler for "/some-path2"
	http.HandleFunc("/some-path2", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request for:", r.URL.Path)
		r.URL.Path = "/" // Modify the path before forwarding (optional)
		proxy2.ServeHTTP(w, r)
	})

	// Additional handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the Reverse Proxy Server!")
	})

	// Start the server
	port := ":8080"
	log.Println("Proxy server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
