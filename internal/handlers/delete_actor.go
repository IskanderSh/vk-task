package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/lib/error/response"
	"github.com/IskanderSh/vk-task/internal/services"
)

func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.DeleteActor"

	log := h.log.With(
		slog.String("op", op))

	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		response.NewErrorResponse(w, log, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	name := r.URL.Query().Get("name")
	if err := h.actorService.DeleteActor(r.Context(), name); err != nil {
		if errors.Is(err, services.ErrActorNotFound) {
			response.NewErrorResponse(w, log, "Actor with this name not found", http.StatusBadRequest, &err)
			return
		}
		response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, &err)
		return
	}
	log.Debug("successfully delete actor")

	w.Write([]byte("successfully delete actor"))
}
