import { Handle, Position, type Node } from "@xyflow/react";
import { FaBiohazard } from "react-icons/fa6";

export function DetectFileHashNode(hash: string): Node[] {
  return [
    {
      id: `file-${hash}`,
      position: {
        x: 0,
        y: 0,
      },
      data: {
        hash: hash,
      },
      type: "detectFileHashNode",
    },
  ];
}

export function DetectFileHashComponent() {
  return (
    <>
      <div>
        <FaBiohazard />
        {/* IP */}
        <Handle
          id="handle-top"
          type="source"
          position={Position.Right}
          className="handle"
        />
        {/* DNS */}
        <Handle
          id="handle-left"
          type="source"
          position={Position.Left}
          className="handle"
        />
      </div>
    </>
  );
}
