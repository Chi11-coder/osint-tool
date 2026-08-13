import dagre from "@dagrejs/dagre";
import { NODE_HEIGHT, NODE_WIDTH } from "../../../../types/common.types";
import type { Edge, Node } from "@xyflow/react";

const DAGRE_OPTIONS = {
  rankdir: "TB",
  nodesep: 80,
  ranksep: 120,
};

export function getLayoutedElements(nodes: Node[], edges: Edge[]) {
  const g = new dagre.graphlib.Graph();
  const visibleNodes = nodes.filter((node) => !node.hidden);
  const visibleNodeIds = new Set(visibleNodes.map((node) => node.id));
  const visibleEdges = edges.filter(
    (edge) =>
      visibleNodeIds.has(edge.source) && visibleNodeIds.has(edge.target),
  );

  g.setDefaultEdgeLabel(() => ({}));
  g.setGraph(DAGRE_OPTIONS);

  visibleNodes.forEach((node) => {
    g.setNode(node.id, {
      width: node.measured?.width ?? NODE_WIDTH,
      height: node.measured?.height ?? NODE_HEIGHT,
    });
  });
  visibleEdges.forEach((edge) => {
    g.setEdge(edge.source, edge.target);
  });

  dagre.layout(g);

  const layoutedNodes = nodes.map((node) => {
    const { x, y } = g.node(node.id);
    const width = node.measured?.width ?? NODE_WIDTH;
    const height = node.measured?.height ?? NODE_HEIGHT;

    return {
      ...node,
      position: {
        x: x - width / 2,
        y: y - height / 2,
      },
    };
  });

  return {
    nodes: layoutedNodes,
    edges,
  };
}
