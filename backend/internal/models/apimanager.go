package models

import "time"

// VirusTotal Domain, IPAddress
type APIManageHostReport struct {
	VirusTotal *VirusTotalHostReports
	CacheTime  time.Time
}

// ファイル調査中央管理
type APIManageFileReport struct {
	VirusTotal *VirusTotalFileReports
}
