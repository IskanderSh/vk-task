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
	"github.com/IskanderSh/vk-task/internal/storage/actor"
	"github.com/IskanderSh/vk-task/internal/storage/film"
	"github.com/IskanderSh/vk-task/internal/storage/user"
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

	actorStorage := actor.NewStorage(db)
	userStorage := user.NewStorage(db)
	filmStorage := film.NewStorage(db)

	// Services
	actorService := services.NewActorService(log, actorStorage)
	filmService := services.NewFilmService(log, filmStorage)
	userService := services.NewUserService(log, userStorage)

	// Handlers
	handler := handlers.NewHandler(log, actorService, filmService, userService)

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
