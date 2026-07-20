package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/car-rental-template/api/internal/config"
	"github.com/car-rental-template/api/internal/server"
)

func main() {
	cfg := config.Load()
	app, closeResources, err := server.New(cfg)
	if err != nil { log.Fatal(err) }
	defer closeResources()
	go app.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Stop(ctx); err != nil { log.Printf("shutdown error: %v", err) }
}
