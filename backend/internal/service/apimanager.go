package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"example.com/security/internal/models"
	"example.com/security/internal/service/hash"
	"example.com/security/internal/service/host"
)

const (
	VirusTotal = "virustotal"
	AbuseIPDB  = "abuseipdb"
	Urlhaus    = "urlhaus"
)

var (
	ServiceTimeout = 15 * time.Second
	ErrQueryParam  = errors.New("invalid query parameter")
)

type VirusTotalReport interface {
	HostReport(ctx context.Context, reportType int, queryByHost string) (*models.VirusTotalHostReports, error)
	FileReport(ctx context.Context, queryByHash string) (*models.VirusTotalFileReports, error)
}

type JsonCache[T any] interface {
	Load(ctx context.Context, v string) (*T, error)
	Save(ctx context.Context, v string, d T) error
}

// 各Servsice層の中央管理サービス
type APIHandlerService struct {
	VirusTotal VirusTotalReport
	JsonCache  JsonCache[models.APIManageHostReport]
}

// IPAddress, Domain 調査
func (s *APIHandlerService) HostReport(ctx context.Context, queryByHost string) (*models.APIManageHostReport, error) {
	reportType, err := host.ReportType(queryByHost)

	if err != nil {
		return nil, fmt.Errorf("%w: param value '%s': %w", ErrQueryParam, queryByHost, err)
	}

	// キャッシュファイルに存在しない, 検索から1日経過した値である場合は API を実行する
	cacheData, err := s.JsonCache.Load(ctx, queryByHost)

	if err != nil {
		return nil, fmt.Errorf("loading host data cache file: %w", err)
	}
	if cacheData != nil && host.SameCacheDate(cacheData.CacheTime) {
		return cacheData, nil
	}
	slog.Info("cache is not found or old")

	apiCtx, cancel := context.WithTimeout(ctx, ServiceTimeout)
	defer cancel()

	report := new(models.APIManageHostReport)
	res, err := s.VirusTotal.HostReport(apiCtx, reportType, queryByHost)

	if err != nil {
		return nil, fmt.Errorf("host report: %w", err)
	}

	// API実行結果
	report.VirusTotal = res
	report.CacheTime = time.Now()

	if err := s.JsonCache.Save(ctx, queryByHost, *report); err != nil {
		slog.Error("can't saving json file", slog.Any("error", err))
	}

	return report, nil
}

func (s *APIHandlerService) FileReport(ctx context.Context, queryByHash string) (*models.APIManageFileReport, error) {
	if !hash.IsValid(queryByHash) {
		return nil, fmt.Errorf("%w: *%s* is not sha-256", ErrQueryParam, queryByHash)
	}

	apiCtx, cancel := context.WithTimeout(ctx, ServiceTimeout)
	defer cancel()

	report := new(models.APIManageFileReport)
	res, err := s.VirusTotal.FileReport(apiCtx, queryByHash)

	if err != nil {
		return nil, fmt.Errorf("file report: %w", err)
	}

	// 検索結果
	report.VirusTotal = res

	return report, nil
}
