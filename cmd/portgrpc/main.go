package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	app, err := BuildDIForApp()
	if err != nil {
		log.Fatalf("Failed to build DI for app: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	log.Print("Starting app")

	go app.server.Run()

	<-ctx.Done()
	stop()

	log.Print("Shutting down the app")
}