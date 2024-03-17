package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/generated/models"
)

type Handler struct {
	log          *slog.Logger
	actorService ActorProvider
	filmService  FilmProvider
}

type ActorProvider interface {
	AddActor(ctx context.Context, actor *models.Actor) error
}

type FilmProvider interface {
}

func NewHandler(
	log *slog.Logger,
	actorProvider ActorProvider,
	filmProvider FilmProvider,
) *Handler {
	return &Handler{
		log:          log,
		actorService: actorProvider,
		filmService:  filmProvider,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/v1/actor/create", h.createActor)

	return router
}
