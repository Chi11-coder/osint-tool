package main

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"example.com/security/internal/handler"
	"example.com/security/internal/models"
	"example.com/security/internal/repository"
	"example.com/security/internal/service"
)

func main() {
	if err := LoadVirusTotalKey(); err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	virusTotalRepo := repository.NewVirusTotalRepository(VirusTotalKey)

	virusTotalSvc := service.VirusTotalService{Store: virusTotalRepo}

	jsonRepo := repository.NewJsonRepository[models.APIManageHostReport]("host-report.json")
	jsonCache := service.JsonService[models.APIManageHostReport]{Store: jsonRepo}

	apiService := service.APIHandlerService{
		VirusTotal: &virusTotalSvc,
		JsonCache:  &jsonCache,
	}
	handle := handler.NewAPIHandler(&apiService)

	distFS, err := fs.Sub(Assets, "dist")

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle("/", handler.NewSPAHandler(distFS))
	mux.HandleFunc("GET /host/report", handle.HostReport)
	mux.HandleFunc("GET /file/report", handle.FileReport)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		slog.ErrorContext(context.Background(), "failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
