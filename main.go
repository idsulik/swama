package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/idsulik/swama/v2/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Handle graceful shutdown with signal handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	err := cmd.Execute(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
