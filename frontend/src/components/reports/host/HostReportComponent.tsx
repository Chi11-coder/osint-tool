import { NoDataComponent } from "../../NoDataComponent";

import "../../../styles/hostReport.css";
import type { HostReport } from "../../../types/host.types";
import { useMemo } from "react";
import { sanitizeIPReports } from "../../../utils/sanitizeIPAddress";
import { DashboardTitle } from "../../../App";

interface HostReportPpops {
  reports: HostReport | null;
  onBack: () => void;
}

export function HostReportCompoent({ reports, onBack }: HostReportPpops) {
  // IPアドレスサニタイズ
  const getSanitizing = useMemo(() => {
    return sanitizeIPReports(reports);
  }, [reports]);

  if (getSanitizing === null) {
    return <NoDataComponent name="Host Report" />;
  }
  return (
    <div>
      <div className="host-report-container">
        <section className="virustotal-summary">
          <h3 className="virustotal-title">
            <DashboardTitle onBack={onBack} />
            <span>VirusTotal Report</span>
          </h3>
          <div className="virustotal-container">
            <div className="virustotal-status malicious">
              <span>{reports?.VirusTotal?.Malicious}</span>
              <span className="status-label">Malicious</span>
            </div>
            <div className="virustotal-status suspicious">
              <span>{reports?.VirusTotal?.Suspicious}</span>
              <span className="status-label">Suspicious</span>
            </div>
            <div className="virustotal-status undetected">
              <span>{reports?.VirusTotal?.Undetected}</span>
              <span className="status-label">Undetected</span>
            </div>
            <div className="virustotal-status harmless">
              <span>{reports?.VirusTotal?.Harmless}</span>
              <span className="status-label">Harmless</span>
            </div>
          </div>
        </section>
      </div>
    </div>
  );
}
