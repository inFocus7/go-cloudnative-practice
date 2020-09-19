package main

import (
	"fabiangonz98/go-cloudnative-practice/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	addr := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", addr)

	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
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
