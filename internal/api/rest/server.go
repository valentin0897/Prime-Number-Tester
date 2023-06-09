package server

import (
	"context"
	handler "primes/internal/api/http/handlers"

	"github.com/labstack/echo"
)

// The Server struct contains an instance of the Echo framework, which is used to handle
// HTTP requests.
type Server struct {
	*echo.Echo
}

// NewServer creates a new instance of the Server
//
// The function also sets up endpoints
func NewServer() *Server {
	e := echo.New()

	server := &Server{e}

	e.POST("/", handler.PrimeNumbersHandler)

	return server
}

// Start launches the Echo HTTP server on port 5000
//
// If something went wrong while starting the server: sends error through channel
func (s *Server) Start(port string, errorCh chan<- error) {
	err := s.Echo.Start(port)
	errorCh <- err
}

// GracefulShutdown is gracefully shutdown the server
func (s *Server) GracefulShutdown(ctx context.Context) {
	s.Echo.Shutdown(ctx)
}
