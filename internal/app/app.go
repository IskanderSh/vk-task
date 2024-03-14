package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/IskanderSh/vk-task/internal/config"
	"github.com/IskanderSh/vk-task/internal/handlers"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(log *slog.Logger, cfg *config.Config) *Server {
	handler := handlers.NewHandler(log)

	router := handler.Routes()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Application.Port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		HTTPServer: httpServer,
	}
}
