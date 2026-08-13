package models

const (
	VirusTotalDomainsURL     = "https://www.virustotal.com/api/v3/domains/"
	VirusTotalIPAddressesURL = "https://www.virustotal.com/api/v3/ip_addresses/"
	VirusTotalFilesURL       = "https://www.virustotal.com/api/v3/files/"
	TypeDomain               = 1
	TypeIP                   = 2
	TypeInvalid              = -1
)

type VirusTotalHostReports struct {
	Host       string
	Malicious  int
	Suspicious int
	Undetected int
	Harmless   int
}

type ReportHostData struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Harmless   int `json:"harmless"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

// ファイル調査
type VirusTotalFileReports struct {
	Hash          string        `json:"Hash"`
	IPTraffic     []IPTraffic   `json:"IPTraffic"`
	DNSLookups    []DNSLookup   `json:"DNSLookups"`
	ProcessesTree []ProcessTree `json:"ProcessesTree"`
}

type ReportFileSummary struct {
	Data struct {
		IPTraffic        []IPTraffic   `json:"ip_traffic"`
		DNSLookups       []DNSLookup   `json:"dns_lookups"`
		ProcessesCreated []string      `json:"processes_created"`
		ProcessesTree    []ProcessTree `json:"processes_tree"`
	} `json:"data"`
}

type ProcessTree struct {
	Name      string        `json:"name"`
	ProcessID string        `json:"process_id"`
	Children  []ProcessTree `json:"children"`
}

type IPTraffic struct {
	DestinationIP   string `json:"destination_ip"`
	DestinationPort int    `json:"destination_port"`
	Protocol        string `json:"transport_layer_protocol"`
}

type DNSLookup struct {
	Hostname    string   `json:"hostname"`
	ResolvedIPs []string `json:"resolved_ips"`
}
