import type { HostReport } from "../types/host.types";
import type { VirusTotalHostReports } from "../types/virustotal.type";

/**
 * ドメインやIPアドレスをxx[.]xx[.]のように表現する
 * @param value
 * @returns string
 */
export function defangDots(value: string): string {
  return value.replace(/(?<!\[)\.(?!\])/g, "[.]");
}

/**
 * API実行結果からドメインやIPアドレス、URLのサニタイズを行う
 * @param reports IPReports
 */
export function sanitizeIPReports(
  reports: HostReport | null,
): HostReport | null {
  if (reports === null) {
    return null;
  }

  const safeVirusTotal = sanitizeVirusTotal(reports?.VirusTotal);
  return {
    ...reports,
    VirusTotal: safeVirusTotal,
  };
}

export function sanitizeVirusTotal(
  report: VirusTotalHostReports | null,
): VirusTotalHostReports | null {
  if (report === null) {
    return null;
  }
  const sanitizeHost = defangDots(report.Host);
  return {
    ...report,
    Host: sanitizeHost,
  };
}
