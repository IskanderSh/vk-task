package response

import (
	"log/slog"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, log *slog.Logger, message string, statusCode int, err *error) {
	if err != nil {
		log.Error((*err).Error())
	} else {
		log.Error(message)
	}
	http.Error(w, message, statusCode)
}
