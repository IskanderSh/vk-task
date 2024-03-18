package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/response"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.Login"

	log := h.log.With(
		slog.String("op", op))

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		response.NewErrorResponse(w, log, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	var input models.UserSignIn

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewErrorResponse(w, log, "Invalid input body", http.StatusBadRequest, &err)
		return
	}
	log.Debug(fmt.Sprintf("successfully decode input body with email: %s", *input.Email))

	token, err := h.userService.Login(r.Context(), &input)
	if err != nil {
		response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, &err)
		return
	}
	log.Debug("successfully authenticate user")

	w.Header().Set("Authorization", token)
	w.Write([]byte(token))
}
