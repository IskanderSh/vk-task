package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/IskanderSh/vk-task/internal/config"
	"github.com/IskanderSh/vk-task/internal/handlers"
	"github.com/IskanderSh/vk-task/internal/services"
	"github.com/IskanderSh/vk-task/internal/storage"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(log *slog.Logger, cfg *config.Config) *Server {
	// Storages
	db, err := storage.NewStorage(&cfg.Storage)
	if err != nil {
		panic(err)
	}

	// Services
	actorService := services.NewActorService(log, db)
	filmService := services.NewFilmService(log, db)

	// Handlers
	handler := handlers.NewHandler(log, actorService, filmService)

	// Init Routes
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
