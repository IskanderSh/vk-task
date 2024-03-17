package handlers

import "net/http"

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}
