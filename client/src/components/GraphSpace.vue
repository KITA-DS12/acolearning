<script setup lang="ts">
import axios from "axios";
import {
  getFullConfigs,
  type Configs,
  type Edges,
  type Layouts,
  type Nodes,
} from "v-network-graph";
import { ref, type Ref } from "vue";

const zoomLevel = 3;

const nodes: Ref<Nodes> = ref({
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
});
let edges: Ref<Edges> = ref({});
const layouts: Ref<Layouts> = ref({
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
});
const configs: Configs = getFullConfigs();

const testSend = async () => {
  await axios
    .post("aco", {
      nodes: nodes.value,
      edges: edges.value,
      layouts: layouts.value,
    })
    .then((res) => {
      edges.value = res.data.edges;
    })
    .catch((error) => {
      console.log(error);
    });
};
</script>

<template>
  <div class="d-flex flex-row justify-center">
    <v-btn @click="testSend" color="indigo">送信</v-btn>
  </div>
  <div class="draw-space">
    <v-network-graph
      ref="graph"
      v-model:zoom-level="zoomLevel"
      v-model:edges="edges"
      v-model:nodes="nodes"
      v-model:layouts="layouts"
      v-model:configs="configs"
    >
    </v-network-graph>
  </div>
</template>
<style>
.draw-space {
  width: 90%;
  height: 50vh;
  border: 1px solid;
  border-radius: 30px;
  border-width: 2px;
  border-color: #7e57c2;
  margin: 3vh auto;
}
</style>
