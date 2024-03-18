package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
}

func (h *Handler) authenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "middleware.AuthenticateAdmin"

		log := h.log.With(slog.String("op", op))

		payload, ok := jwtPayloadFromRequest(r, log)
		if !ok {
			return
		}

		log.Info(fmt.Sprintf("get payload with params: sub - %s, role - %s", payload["sub"].(string), payload["role"].(string)))
	})
}

func AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func jwtPayloadFromRequest(r *http.Request, log *slog.Logger) (jwt.MapClaims, bool) {
	jwtToken, ok := r.Context().Value("Authenticate").(*jwt.Token)
	if !ok {
		log.Error("wrong type of JWT token")
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}
