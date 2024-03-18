package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/IskanderSh/vk-task/internal/lib/error/response"
	"github.com/IskanderSh/vk-task/internal/services"
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
			response.NewErrorResponse(w, log, "Unauthorized user", http.StatusUnauthorized, nil)
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
	headerValue := r.Header.Get("Authorization")

	jwtToken, err := jwt.ParseWithClaims(headerValue, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(services.SigningKey), nil
	})
	if err != nil {
		log.Error("couldn't get jwt token from header")
		log.Error(err.Error())
	}
	log.Debug("successfully get jwt token from header")

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}
