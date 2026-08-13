package host

import (
	"errors"
	"net"
	"regexp"

	"example.com/security/internal/models"
)

var (
	ErrInvalidParamHost = errors.New("not match domain or ip")
	domainRegex         = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
)

func ReportType(host string) (int, error) {
	if domainRegex.MatchString(host) {
		return models.TypeDomain, nil
	}

	ip := net.ParseIP(host)
	if ip != nil {
		return models.TypeIP, nil
	}

	return models.TypeInvalid, ErrInvalidParamHost
}
