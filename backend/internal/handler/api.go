//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=handler
package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"example.com/security/internal/models"
)

type APIManagerService interface {
	HostReport(ctx context.Context, host string) (*models.APIManageHostReport, error)
	FileReport(ctx context.Context, hash string) (*models.APIManageFileReport, error)
}

type Handler struct {
	APIService APIManagerService
}

func NewAPIHandler(s APIManagerService) *Handler {
	return &Handler{APIService: s}
}

func (h *Handler) HostReport(w http.ResponseWriter, r *http.Request) {
	queryByHost := r.URL.Query().Get("host")

	res, err := h.APIService.HostReport(r.Context(), queryByHost)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed host report", slog.Any("error", err))
		ResponseWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode json response", slog.Any("error", err))
		return
	}
}

func (h *Handler) FileReport(w http.ResponseWriter, r *http.Request) {
	queryByHash := r.URL.Query().Get("hash")

	res, err := h.APIService.FileReport(r.Context(), queryByHash)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed file report", slog.Any("error", err))
		ResponseWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode json response", slog.Any("error", err))
		return
	}
}
