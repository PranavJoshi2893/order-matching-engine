package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PranavJoshi2893/order-matching-engine/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Port,
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	errChan := make(chan error, 1)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return fmt.Errorf("failed to start server: %w", err)
	case <-quit:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		log.Println("shutting down....")
		if err := s.httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf("forced shutdown: %w", err)
		}
		log.Println("shutdown")

		return nil
	}

}
