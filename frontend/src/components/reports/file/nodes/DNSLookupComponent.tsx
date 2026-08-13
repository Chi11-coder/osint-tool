import { Handle, Position, type Node } from "@xyflow/react";
import { FaServer } from "react-icons/fa6";
import type { DNSLookups } from "../../../../types/virustotal.type";

interface ParentDNSLookupProps {
  data: {
    label: string;
  };
}
interface ChildDNSLookupProps {
  data: {
    hostName: string;
    resolvedIPs: string[] | undefined;
  };
}

export function ParentDNSLookupNode(): Node[] {
  return [
    {
      id: "dns-lookup-node",
      position: {
        x: 0,
        y: 0,
      },
      data: {
        label: "DNS",
      },
      type: "parentDNSLookupNode",
    },
  ];
}

export function ParentDNSLookupComponent({ data }: ParentDNSLookupProps) {
  return (
    <div>
      <FaServer />
      <span>{data.label}</span>
      <Handle
        id="handle-left"
        type="source"
        position={Position.Bottom}
        className="handle"
      />
      <Handle type="target" position={Position.Top} className="handle" />
    </div>
  );
}

export function ChildDNSLookupNode(dnsLookups: Array<DNSLookups>): Node[] {
  if (!dnsLookups) {
    return [
      {
        id: "child-dns-00",
        position: { x: 0, y: 0 },
        data: {
          hostName: "",
          resolvedIPs: undefined,
        },
        type: "childDNSTrafficNode",
      },
    ];
  }

  return dnsLookups.map((d, idx) => {
    return {
      id: `child-dns-${idx}`,
      position: {
        x: 0,
        y: 0,
      },
      data: {
        hostName: d.hostname,
        group: "child-dns-lookup",
      },
      type: "childDNSTrafficNode",
    };
  });
}

export function ChildDNSLookupComponent({ data }: ChildDNSLookupProps) {
  return (
    <div>
      <div>
        <div>
          <span>{data.hostName}</span>
        </div>
      </div>
      <Handle type="target" position={Position.Top} className="handle" />
    </div>
  );
}
