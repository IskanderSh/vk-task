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

func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.CreateFilm"

	log := h.log.With(
		slog.String("op", op))

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		response.NewErrorResponse(w, log, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var input models.Film

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewErrorResponse(w, log, "Invalid input body", http.StatusBadRequest, &err)
		return
	}
	log.Debug(fmt.Sprintf("successfully decode input body with name: %s", *input.Name))

	if err := h.filmService.AddFilm(r.Context(), &input); err != nil {
		if errors.Is(err, services.ErrDuplicateName) {
			response.NewErrorResponse(w, log, "Invalid film's name", http.StatusBadRequest, &err)
			return
		}
		response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, &err)
		return
	}
	log.Debug("successfully create new film")

	w.Write([]byte("successfully created new film"))
}
