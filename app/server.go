package app

import (
	"fabiangonz98/go-cloudnative-practice/util/logger"
	"github.com/jinzhu/gorm"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{logger, db}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
