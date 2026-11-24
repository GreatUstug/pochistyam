package app

import "pochistyam/internal/db"

type App struct {
	DB *db.DB
}

func NewApp(database *db.DB) *App {
	return &App{DB: database}
}
