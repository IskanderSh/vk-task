package handlers

import (
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/lib/error/response"
	"github.com/IskanderSh/vk-task/internal/services"
	"github.com/golang-jwt/jwt/v5"
)

var (
	adminRole = "admin"
	userRole  = "user"
)

func (h *Handler) authenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "middleware.AuthenticateAdmin"

		log := h.log.With(slog.String("op", op))

		payload, ok := jwtPayloadFromRequest(r, log)
		if !ok {
			response.NewErrorResponse(w, log, "Unauthorized user", http.StatusUnauthorized, nil)
			return
		}

		role, ok := payload["role"].(string)
		if !ok {
			response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, nil)
			return
		}

		if role != adminRole {
			response.NewErrorResponse(w, log, "Don't have enough permission", http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "middleware.AuthenticateAdmin"

		log := h.log.With(slog.String("op", op))

		payload, ok := jwtPayloadFromRequest(r, log)
		if !ok {
			response.NewErrorResponse(w, log, "Unauthorized user", http.StatusUnauthorized, nil)
			return
		}

		role, ok := payload["role"].(string)
		if !ok {
			response.NewErrorResponse(w, log, "Internal error", http.StatusInternalServerError, nil)
			return
		}

		if role != adminRole || role != userRole {
			response.NewErrorResponse(w, log, "Don't have enough permission", http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func jwtPayloadFromRequest(r *http.Request, log *slog.Logger) (jwt.MapClaims, bool) {
	headerValue := r.Header.Get("Authorization")

	jwtToken, err := jwt.ParseWithClaims(headerValue, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(services.SigningKey), nil
	})
	if err != nil {
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
