package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/generated/models"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.UserSignUp

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	if err := h.userService.AddUser(r.Context(), &input); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("successfully created new user"))
}
