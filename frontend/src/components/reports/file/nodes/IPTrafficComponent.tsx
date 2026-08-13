import { Handle, Position, type Node } from "@xyflow/react";
import { FaLocationDot } from "react-icons/fa6";
import type { IPTraffics } from "../../../../types/virustotal.type";

interface ParentIPTrafficProps {
  data: {
    label: string;
  };
}
interface ChildIPTrafficProps {
  data: {
    ip: string;
    port: number | undefined;
  };
}

export function ParentIPTrafficNode(): Node[] {
  return [
    {
      id: "ip-traffic-node",
      position: {
        x: 0,
        y: 0,
      },
      data: {
        label: "IP",
      },
      type: "parentIPTrafficNode",
    },
  ];
}

export function ParentIPTrafficComponent({ data }: ParentIPTrafficProps) {
  return (
    <div>
      <FaLocationDot />
      <span>{data.label}</span>
      <Handle
        id="handle-top"
        type="source"
        position={Position.Bottom}
        className="handle"
      />
      <Handle type="target" position={Position.Top} className="handle" />
    </div>
  );
}

export function ChildIPTrafficNode(ipTraffic: Array<IPTraffics>): Node[] {
  if (ipTraffic.length === 0) {
    return [
      {
        id: "child-ip-00",
        position: {
          x: 0,
          y: 0,
        },
        data: {
          ip: "",
          port: undefined,
        },
        type: "childIPTrafficNode",
      },
    ];
  }

  return ipTraffic.map((ip, idx) => {
    return {
      id: `child-ip-${idx}`,
      position: {
        x: 0,
        y: 0,
      },
      data: {
        ip: ip.destination_ip,
        port: ip.destination_port,
        group: "child-ip-traffic",
      },
      type: "childIPTrafficNode",
    };
  });
}

export function ChildIPTrafficComponent({ data }: ChildIPTrafficProps) {
  return (
    <div>
      <span>
        {data.ip}:{data.port}
      </span>
      <Handle type="target" position={Position.Top} className="handle" />
    </div>
  );
}
