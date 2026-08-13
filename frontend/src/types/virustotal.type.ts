/** IP,ドメイン調査 **/
export type VirusTotalHostReports = {
  Host: string;
  Malicious: number;
  Suspicious: number;
  Undetected: number;
  Harmless: number;
};

/** ファイル調査 **/
export interface VirusTotalFileReports {
  Hash: string;
  IPTraffic: Array<IPTraffics>;
  DNSLookups: Array<DNSLookups>;
  ProcessesTree: Array<ProcessesTree>;
}

export interface IPTraffics {
  destination_ip: string;
  destination_port: number;
  transport_layer_protocol: string;
}

export interface DNSLookups {
  hostname: string;
  resolved_ips: string[];
}

export interface ProcessesTree {
  name: string;
  process_id: string;
  children: Array<ProcessesTree>;
}
