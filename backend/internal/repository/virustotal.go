//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=repository
package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"example.com/security/internal/models"
	"example.com/security/internal/repository/processtree"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrUnauthorized          = errors.New("unauthenticate virus total api key")
	ErrForbidden             = errors.New("forbidden virus total api key")
	ErrNotFound              = errors.New("not found")
	ErrRateLimit             = errors.New("too many requests")
	ErrGatewayTimeout        = errors.New("gateway timeout")
	ErrHTTPClient            = errors.New("client error")
	ErrHTTPServer            = errors.New("server error")
	ErrUnsupportedReportType = errors.New("unsupported report type")
)

type VirusTotalRepository interface {
	HostReport(ctx context.Context, reportType int, queryByHost string) (*models.VirusTotalHostReports, error)
	FileReport(ctx context.Context, hash string) (*models.VirusTotalFileReports, error)
	get(ctx context.Context, endpoint string) (*http.Response, error)
}

type virusTotalRepo struct {
	apiKey string
	client *http.Client
}

// コンストラクタ
func NewVirusTotalRepository(key string) VirusTotalRepository {
	return &virusTotalRepo{
		apiKey: key,
		client: &http.Client{},
	}
}

func (r *virusTotalRepo) HostReport(ctx context.Context, reportType int, host string) (*models.VirusTotalHostReports, error) {
	url, err := queryHostURL(reportType, host)

	if err != nil {
		return nil, fmt.Errorf("unknown report type: %w", err)
	}

	res, err := r.get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("http request error: %w", err)
	}
	defer res.Body.Close()

	var report models.ReportHostData
	if err := json.NewDecoder(res.Body).Decode(&report); err != nil {
		return nil, fmt.Errorf("json file decode error: %w", err)
	}

	data := report.Data.Attributes.LastAnalysisStats

	return &models.VirusTotalHostReports{
		Host:       host,
		Malicious:  data.Malicious,
		Suspicious: data.Suspicious,
		Undetected: data.Undetected,
		Harmless:   data.Harmless,
	}, nil
}

func queryHostURL(reportType int, queryByHost string) (string, error) {
	switch reportType {
	case models.TypeDomain:
		return models.VirusTotalDomainsURL + url.PathEscape(queryByHost), nil
	case models.TypeIP:
		return models.VirusTotalIPAddressesURL + url.PathEscape(queryByHost), nil
	default:
		return "", fmt.Errorf("%w: %d", ErrUnsupportedReportType, reportType)
	}
}

func (r *virusTotalRepo) FileReport(ctx context.Context, queryByHash string) (*models.VirusTotalFileReports, error) {
	url := models.VirusTotalFilesURL + queryByHash + "/behaviour_summary"
	res, err := r.get(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("http request error: %w", err)
	}
	defer res.Body.Close()

	report := new(models.ReportFileSummary)
	if err := json.NewDecoder(res.Body).Decode(&report); err != nil {
		return nil, fmt.Errorf("json file decode error: %w", err)
	}

	ipTraffics := mapSlice(report.Data.IPTraffic, func(t models.IPTraffic) models.IPTraffic {
		return models.IPTraffic{
			DestinationIP:   t.DestinationIP,
			DestinationPort: t.DestinationPort,
			Protocol:        t.Protocol,
		}
	})

	dnsLookups := mapSlice(report.Data.DNSLookups, func(d models.DNSLookup) models.DNSLookup {
		reslovedIPs := d.ResolvedIPs
		if reslovedIPs == nil {
			reslovedIPs = []string{}
		}

		return models.DNSLookup{
			Hostname:    d.Hostname,
			ResolvedIPs: reslovedIPs,
		}
	})

	prunedTree := processtree.Prune(report.Data.ProcessesCreated, report.Data.ProcessesTree)
	processTree := processtree.Build(prunedTree)
	return &models.VirusTotalFileReports{
		Hash:          queryByHash,
		IPTraffic:     ipTraffics,
		DNSLookups:    dnsLookups,
		ProcessesTree: processTree,
	}, nil
}

func mapSlice[S, T any](src []S, fn func(S) T) []T {
	result := make([]T, len(src))
	for i, v := range src {
		result[i] = fn(v)
	}
	return result
}

// Common Methtod Get
func (r *virusTotalRepo) get(ctx context.Context, endpoing string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoing, nil)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	req.Header.Set("x-apikey", r.apiKey)
	req.Header.Set("accept", "application/json")

	res, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, responseError(res.StatusCode)
	}

	return res, nil
}

func responseError(statusCode int) error {
	switch statusCode {
	case http.StatusBadRequest:
		return fmt.Errorf("%w: status code: %d", ErrBadRequest, statusCode)
	case http.StatusUnauthorized:
		return fmt.Errorf("%w: status code: %d", ErrUnauthorized, statusCode)
	case http.StatusForbidden:
		return fmt.Errorf("%w: status code: %d", ErrForbidden, statusCode)
	case http.StatusNotFound:
		return fmt.Errorf("%w: status code: %d", ErrNotFound, statusCode)
	case http.StatusTooManyRequests:
		return fmt.Errorf("%w: status code: %d", ErrRateLimit, statusCode)
	case http.StatusGatewayTimeout:
		return fmt.Errorf("%w: status code: %d", ErrGatewayTimeout, statusCode)
	default:
		if statusCode >= http.StatusInternalServerError {
			return fmt.Errorf("%w: status code: %d", ErrHTTPServer, statusCode)
		}
		return fmt.Errorf("%w: status code: %d", ErrHTTPClient, statusCode)
	}
}
