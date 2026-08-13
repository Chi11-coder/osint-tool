import type { VirusTotalHostReports } from "./virustotal.type";

export interface HostReport {
  VirusTotal: VirusTotalHostReports | null;
  AddCacheTime: number;
}
