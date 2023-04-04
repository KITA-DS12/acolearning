<script setup lang="ts">
import { useStore } from "@/store";
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
const store = useStore();

const nodes: Ref<Nodes> = ref(store.nodes);
const edges: Ref<Edges> = ref(store.edges);
const layouts: Ref<Layouts> = ref(store.layouts);
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
