package helpers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetStatusCode of http status
func GetStatusCode(err error) int {
	if err == nil || err == ErrNotFound {
		return http.StatusOK
	}

	log.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
