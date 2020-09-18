package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	log.Println("Starting server :8080")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed! >> %v", err)
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello World!"); err != nil {
		log.Fatalf("Error contacting handler... >> %v", err)
	}
}