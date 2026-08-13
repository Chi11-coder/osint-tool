import type { Edge } from "@xyflow/react";
import type { DNSLookups, IPTraffics } from "../../../../types/virustotal.type";

export function ChildIPTrafficNodeToEdge(
  ipTraffics?: Array<IPTraffics>,
): Edge[] {
  const safeIPTraffics = ipTraffics ?? [];

  return safeIPTraffics.map((_, idx) => {
    return {
      id: `child-ip-traffic-${idx}`,
      source: `ip-traffic-node`,
      target: `child-ip-${idx}`,
      sourceHandle: "handle-top",
      data: {
        group: "child-ip-traffic",
      },
      markerEnd: { type: "arrowclosed" },
      animated: true,
    };
  });
}

export function ChildDNSLookupNodeToEdge(
  dnsLookups?: Array<DNSLookups>,
): Edge[] {
  const safeDNSLookups = dnsLookups ?? [];

  return safeDNSLookups.map((_, idx) => {
    return {
      id: `child-dns-lookup-${idx}`,
      source: "dns-lookup-node",
      target: `child-dns-${idx}`,
      sourceHandle: "handle-left",
      data: {
        group: "child-dns-lookup",
      },
      markerEnd: { type: "arrowclosed" },
      animated: true,
    };
  });
}
