package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	server "primes/internal/api/rest"
	"syscall"
)

var (
	port = flag.String("port", ":5001", "server port")
)

func main() {
	flag.Parse()

	srv := server.NewServer()

	ctx := context.TODO()

	done := make(chan error)

	go func() {
		srv.Start(*port, done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	handleShutdown(srv, quit, done)

	srv.GracefulShutdown(ctx)
}

func handleShutdown(srv *server.Server, quit <-chan os.Signal, done <-chan error) {
	select {
	case q := <-quit:
		srv.Echo.Logger.Errorf("program was interrupted: %v", q)
		return

	case serverError := <-done:
		srv.Echo.Logger.Errorf("something went wrong while starting the server: %v", serverError)
		return
	}
}
