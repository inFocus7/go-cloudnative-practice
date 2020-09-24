package main

import (
	dbConn "fabiangonz98/go-cloudnative-practice/adapter/gorm"
	"fabiangonz98/go-cloudnative-practice/app"
	"fabiangonz98/go-cloudnative-practice/app/router"
	"fabiangonz98/go-cloudnative-practice/config"
	"fabiangonz98/go-cloudnative-practice/util/logger"
	"fmt"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	logger := logger.New(appConf.Debug, false)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		db.LogMode(true)
	}

	application := app.New(logger, db)
	appRouter := router.New(application)

	addr := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", addr)

	s := &http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
