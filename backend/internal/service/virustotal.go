package service

import (
	"context"

	"example.com/security/internal/models"
	"example.com/security/internal/repository"
)

type VirusTotalService struct {
	Store repository.VirusTotalRepository
}

func (s *VirusTotalService) HostReport(ctx context.Context, reportType int, queryByHost string) (*models.VirusTotalHostReports, error) {
	return s.Store.HostReport(ctx, reportType, queryByHost)
}

func (s *VirusTotalService) FileReport(ctx context.Context, queryByHash string) (*models.VirusTotalFileReports, error) {
	return s.Store.FileReport(ctx, queryByHash)
}
