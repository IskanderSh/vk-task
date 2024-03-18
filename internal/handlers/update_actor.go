package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/response"
	"github.com/IskanderSh/vk-task/internal/services"
)

func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.createActor"

	log := h.log.With(
		slog.String("op", op))

	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		response.NewErrorResponse(w, log, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var input models.UpdateActor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewErrorResponse(w, log, "Invalid input body", http.StatusBadRequest, &err)
		return
	}
	log.Debug(fmt.Sprintf("successfully decode input body with name: %s", input.Name))

	if err := h.actorService.UpdateActor(r.Context(), &input); err != nil {
		if errors.Is(err, services.ErrActorNotFound) {
			response.NewErrorResponse(w, log, "Actor with this name not found", http.StatusBadRequest, &err)
			return
		}
		response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, &err)
		return
	}
	log.Debug("successfully update actor")

	w.Write([]byte("successfully updated actor"))
}
