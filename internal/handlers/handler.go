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
	UpdateActor(ctx context.Context, input *models.UpdateActor) error
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

	router.Handle("/api/v1/actor/create", h.authenticateAdmin(http.HandlerFunc(h.CreateActor)))
	router.Handle("/api/v1/actor/update", h.authenticateAdmin(http.HandlerFunc(h.UpdateActor)))
	router.Handle("/api/v1/actor/delete/:id", h.authenticateAdmin(http.HandlerFunc(h.DeleteActor)))

	router.Handle("/api/v1/film/create", h.authenticateAdmin(http.HandlerFunc(h.CreateFilm)))
	router.Handle("/api/v1/film/update", h.authenticateAdmin(http.HandlerFunc(h.UpdateFilm)))
	router.Handle("/api/v1/film/delete/:id", h.authenticateAdmin(http.HandlerFunc(h.DeleteFilm)))

	router.Handle("/api/v1/films/:sortby", h.authenticateUser(http.HandlerFunc(h.Films)))
	router.Handle("/api/v1/films/:name", h.authenticateUser(http.HandlerFunc(h.FilmsByName)))
	router.Handle("/api/v1/films/actor/:name", h.authenticateUser(http.HandlerFunc(h.FilmsByActor)))
	router.Handle("/api/v1/actors", h.authenticateUser(http.HandlerFunc(h.ActorsWithFilms)))

	return router
}
