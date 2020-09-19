package app

import "fabiangonz98/go-cloudnative-practice/util/logger"

type App struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *App {
	return &App{logger}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
