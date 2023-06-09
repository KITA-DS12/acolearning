import { defineStore } from "pinia";
import type { Nodes, Edges, Layouts } from "v-network-graph";

interface GraphState {
  zoomLevel: number;
  selectedNodes: Array<string>;
  nodes: Nodes;
  edges: Edges;
  layouts: Layouts;
}

export const useStore = defineStore({
  id: "graph",
  state: (): GraphState => ({
    zoomLevel: 5,
    nodes: {
      node0: { name: "Node 0" },
      node1: { name: "Node 1" },
      node2: { name: "Node 2" },
      node3: { name: "Node 3" },
      node4: { name: "Node 4" },
      node5: { name: "Node 5" },
      node6: { name: "Node 6" },
      node7: { name: "Node 7" },
      node8: { name: "Node 8" },
      node9: { name: "Node 9" },
      node10: { name: "Node 10" },
      node11: { name: "Node 11" },
      node12: { name: "Node 12" },
      node13: { name: "Node 13" },
      node14: { name: "Node 14" },
      node15: { name: "Node 15" },
      node16: { name: "Node 16" },
      node17: { name: "Node 17" },
      node18: { name: "Node 18" },
      node19: { name: "Node 19" },
    },
    edges: {},
    layouts: {
      nodes: {
        node0: { x: 0, y: 0 },
        node1: { x: 20, y: 10 },
        node2: { x: 40, y: 10 },
        node3: { x: 60, y: 30 },
        node4: { x: 80, y: 40 },
        node5: { x: 100, y: 40 },
        node6: { x: 120, y: 30 },
        node7: { x: 140, y: 10 },
        node8: { x: 160, y: 10 },
        node9: { x: 180, y: 0 },
        node10: { x: 0, y: 40 },
        node11: { x: 20, y: 30 },
        node12: { x: 40, y: 30 },
        node13: { x: 60, y: 10 },
        node14: { x: 80, y: 0 },
        node15: { x: 100, y: 0 },
        node16: { x: 120, y: 10 },
        node17: { x: 140, y: 30 },
        node18: { x: 160, y: 30 },
        node19: { x: 180, y: 40 },
      },
    },
    selectedNodes: [],
  }),
  actions: {
    setLevel(newValue: number) {
      this.zoomLevel = newValue;
    },
    setNodes(newArray: Nodes) {
      this.nodes = newArray;
    },
    setSelectedNodes(newArray: Array<string>) {
      this.selectedNodes = newArray;
    },
    addNode(id: string, name: string) {
      const nodes = this.nodes;
      nodes[id] = { name: name };
      this.nodes = nodes;
    },
    removeNode(id: string) {
      delete this.nodes[id];
    },
  },
});
