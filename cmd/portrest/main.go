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
		log.Fatalf("Failed to build DI for ports rest app: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	log.Print("Starting port rest app")

	go app.router.Run()

	<-ctx.Done()
	stop()

	log.Print("Shutting down ports rest app")
}
