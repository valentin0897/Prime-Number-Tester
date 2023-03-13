package main

import (
	"context"
	"os"
	"os/signal"
	server "primes/internal/api/rest"
	"syscall"
)

func main() {
	srv := server.NewServer()

	ctx := context.TODO()

	go func() {
		srv.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

shutdown:
	for {
		select {
		case v := <-quit:
			srv.Echo.Logger.Errorf("program was interrupted: %v", v)
			break shutdown
		case done := <-ctx.Done():
			srv.Echo.Logger.Errorf("context was cancelled: %v", done)
			break shutdown
		}
	}

	srv.GracefulShutdown(ctx)
}
