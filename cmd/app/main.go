package main

import (
	"fabiangonz98/go-cloudnative-practice/app/router"
	"fabiangonz98/go-cloudnative-practice/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	addr := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", addr)

	s := &http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed! >> %v", err)
	}
}
