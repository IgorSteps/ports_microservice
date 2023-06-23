package main

import (
	"ports_microservice/internal/drivers/rest"
)

type App struct {
	router *rest.Router
}

func NewApp(r *rest.Router) *App {
	return &App{
		router: r,
	}
}
