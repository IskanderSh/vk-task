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
	userService  UserProvider
}

type UserProvider interface {
	AddUser(ctx context.Context, input *models.UserSignUp) error
	Login(ctx context.Context, input *models.UserSignIn) (string, error)
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
	userProvider UserProvider,
) *Handler {
	return &Handler{
		log:          log,
		actorService: actorProvider,
		filmService:  filmProvider,
		userService:  userProvider,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/auth/sign-up", h.Register)
	router.HandleFunc("/auth/sign-in", h.Login)

	router.Handle("/api/v1/actor/create", h.authenticateAdmin(http.HandlerFunc(h.createActor)))

	return router
}
