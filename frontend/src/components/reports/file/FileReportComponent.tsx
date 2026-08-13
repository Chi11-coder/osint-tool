import {
  addEdge,
  Background,
  BackgroundVariant,
  ReactFlow,
  ReactFlowProvider,
  useEdgesState,
  useNodesInitialized,
  useNodesState,
  useReactFlow,
  type Connection,
  type Edge,
  type Node,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import type { FileReport } from "../../../types/file.types";
import React, { useCallback, useEffect, useState } from "react";
import { DashboardTitle } from "../../../App";
import {
  ChildIPTrafficComponent,
  ChildIPTrafficNode,
  ParentIPTrafficComponent,
  ParentIPTrafficNode,
} from "./nodes/IPTrafficComponent";
import { ParentToChildEdge } from "./edges/parentToChildEdge";
import {
  ChildDNSLookupNodeToEdge,
  ChildIPTrafficNodeToEdge,
} from "./edges/childNodeToEdge";

import "../../../styles/fileReport.css";
import {
  ChildDNSLookupComponent,
  ChildDNSLookupNode,
  ParentDNSLookupComponent,
  ParentDNSLookupNode,
} from "./nodes/DNSLookupComponent";
import {
  DetectFileHashComponent,
  DetectFileHashNode,
} from "./nodes/DetectFileHashComponent";
import { getLayoutedElements } from "./layout/getLayoutedElements";

interface FileReportProps {
  reports: FileReport;
  onBack: () => void;
}

// カスタムノード
const nodeTypes = {
  detectFileHashNode: DetectFileHashComponent,
  parentIPTrafficNode: ParentIPTrafficComponent,
  childIPTrafficNode: ChildIPTrafficComponent,
  parentDNSLookupNode: ParentDNSLookupComponent,
  childDNSTrafficNode: ChildDNSLookupComponent,
};

export function FileReportFlowProvider({ reports, onBack }: FileReportProps) {
  return (
    <ReactFlowProvider>
      <FileReportComponent reports={reports} onBack={onBack} />
    </ReactFlowProvider>
  );
}

export function FileReportComponent({ reports, onBack }: FileReportProps) {
  const [nodes, setNodes, onNodeChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgeChange] = useEdgesState<Edge>([]);

  const { fitView } = useReactFlow();
  const nodesInitialized = useNodesInitialized();

  // IPTraffic 子ノード
  const [isChildIPTraffic, setChildIPTraffic] = useState<boolean>(false);
  // DNSLookup 子ノード
  const [isChildDNSLookup, setChildDNSLookup] = useState<boolean>(false);

  const onConnect = useCallback(
    (changes: Connection) => {
      setEdges((edgeSnap) => addEdge(changes, edgeSnap));
    },
    [setEdges],
  );

  // ノードクリックで子ノードを展開
  // TODO: プロセスツリー展開は未実装
  const handleParentNodeClick = useCallback(
    (_event: React.MouseEvent, node: Node) => {
      if (node.id === "ip-traffic-node") {
        setChildIPTraffic((event) => !event);
      }

      if (node.id === "dns-lookup-node") {
        setChildDNSLookup((event) => !event);
      }
    },
    [],
  );

  // Node, Edge 定義
  useEffect(() => {
    // 初期ノード
    const nodeArray: Node[] = [
      ...DetectFileHashNode(reports?.VirusTotal?.Hash),
      ...ParentIPTrafficNode(),
      ...ParentDNSLookupNode(),
    ];
    const edgeArray: Edge[] = [...ParentToChildEdge(reports?.VirusTotal?.Hash)];

    const ipTraffics = reports?.VirusTotal?.IPTraffic ?? [];
    const dnsLookups = reports?.VirusTotal?.DNSLookups ?? [];

    if (ipTraffics.length !== 0) {
      nodeArray.push(...ChildIPTrafficNode(ipTraffics));
      edgeArray.push(...ChildIPTrafficNodeToEdge(ipTraffics));
    }

    if (dnsLookups.length !== 0) {
      nodeArray.push(...ChildDNSLookupNode(dnsLookups));
      edgeArray.push(...ChildDNSLookupNodeToEdge(dnsLookups));
    }

    const { nodes: layoutedNodes, edges: layoutedEdges } = getLayoutedElements(
      nodeArray,
      edgeArray,
    );

    setNodes(layoutedNodes);
    setEdges(layoutedEdges);
  }, [reports, setNodes, setEdges]);

  useEffect(() => {
    if (nodesInitialized) {
      fitView({ duration: 400, padding: 0.2 });
    }
    setNodes((prev) =>
      prev.map((node) => {
        switch (node.data?.group) {
          case "child-ip-traffic":
            return { ...node, hidden: !isChildIPTraffic };
          case "child-dns-lookup":
            return { ...node, hidden: !isChildDNSLookup };
          default:
            return node;
        }
      }),
    );
    setEdges((prev) =>
      prev.map((edge) => {
        switch (edge.data?.group) {
          case "child-ip-traffic":
            return { ...edge, hidden: !isChildIPTraffic };
          case "child-dns-lookup":
            return { ...edge, hidden: !isChildDNSLookup };
          default:
            return edge;
        }
      }),
    );
  }, [
    isChildIPTraffic,
    isChildDNSLookup,
    nodesInitialized,
    fitView,
    setNodes,
    setEdges,
  ]);

  return (
    <div className="file-report-container">
      <header className="file-report-header">
        <h1 className="file-report-title">
          <DashboardTitle onBack={onBack} />
          <span>VirusTotal ファイル調査</span>
        </h1>
      </header>

      <div className="file-report-flow">
        <ReactFlow
          className="file-report-flow-canvas"
          nodes={nodes}
          edges={edges}
          nodeTypes={nodeTypes}
          onNodesChange={onNodeChange}
          onEdgesChange={onEdgeChange}
          onConnect={onConnect}
          onNodeClick={handleParentNodeClick}
          preventScrolling={false}
          deleteKeyCode={null}
          fitView
        >
          <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
        </ReactFlow>
      </div>
    </div>
  );
}
