import type { Edge } from "@xyflow/react";

export const ParentToChildEdge = (hash: string): Edge[] => {
  return [
    {
      id: "parent-ip-traffic-edge",
      source: `file-${hash}`,
      target: `ip-traffic-node`,
      sourceHandle: "handle-top",
      markerEnd: { type: "arrowclosed" },
      animated: true,
    },
    {
      id: "parent-dns-lookup-edge",
      source: `file-${hash}`,
      target: `dns-lookup-node`,
      sourceHandle: "handle-left",
      markerEnd: { type: "arrowclosed" },
      animated: true,
    },
  ];
};
