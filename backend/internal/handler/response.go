package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"example.com/security/internal/repository"
	"example.com/security/internal/service"
)

type ErrMsg struct {
	Code int
	Msg  string
}

func ResponseWithError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	// Timeout
	case errors.Is(err, context.DeadlineExceeded):
		responseWriter(w, http.StatusGatewayTimeout, "upstream api timeout")
	// Query Param Error
	case errors.Is(err, service.ErrQueryParam):
		responseWriter(w, http.StatusBadRequest, "invalid query param value")
	// Bad Request
	case errors.Is(err, repository.ErrBadRequest):
		responseWriter(w, http.StatusBadRequest, "bad request")
	// Unauthorized
	case errors.Is(err, repository.ErrUnauthorized):
		responseWriter(w, http.StatusUnauthorized, "unauthorized api key")
	// Forbidden
	case errors.Is(err, repository.ErrForbidden):
		responseWriter(w, http.StatusForbidden, "your virustotal api key is forbidden")
	// Not Found
	case errors.Is(err, repository.ErrNotFound):
		responseWriter(w, http.StatusNotFound, "not found")
	// Rate Limit
	case errors.Is(err, repository.ErrRateLimit):
		responseWriter(w, http.StatusTooManyRequests, "api rate limit")
	// Gateway Timeout
	case errors.Is(err, repository.ErrGatewayTimeout):
		responseWriter(w, http.StatusGatewayTimeout, "gatewaty timeout")
	// Others 4xx error
	case errors.Is(err, repository.ErrHTTPClient):
		responseWriter(w, http.StatusBadGateway, "upstream api rejected the request")
	// Others 5xx error
	case errors.Is(err, repository.ErrHTTPServer):
		responseWriter(w, http.StatusBadGateway, "upstream api is unavailable")
	// Internal Server Error
	default:
		responseWriter(w, http.StatusInternalServerError, "internal server error")
	}
}

func responseWriter(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(ErrMsg{
		Code: code,
		Msg:  msg,
	}); err != nil {
		slog.Error("encoding error")
	}
}
