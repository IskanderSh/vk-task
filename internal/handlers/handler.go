package handlers

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/v1/actor/create", h.createActor)

	return router
}
