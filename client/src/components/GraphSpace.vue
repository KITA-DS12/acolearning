<script setup lang="ts">
import { useStore } from "@/store";
import axios from "axios";
import { getFullConfigs, type Configs } from "v-network-graph";
import { computed, ref } from "vue";

const store = useStore();

const zoomLevel = computed(() => store.zoomLevel);
const nodes = computed(() => store.nodes);
const edges = ref(store.edges);
const layouts = computed(() => store.layouts);
const configs: Configs = getFullConfigs();
configs.node.selectable = true;

const selectNodes = computed({
  get: () => store.selectedNodes,
  set: (newArray) => store.setSelectedNodes(newArray),
});

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
  <div class="draw-space">
    <v-network-graph
      ref="graph"
      v-model:zoom-level="zoomLevel"
      v-model:edges="edges"
      v-model:nodes="nodes"
      v-model:layouts="layouts"
      v-model:configs="configs"
      v-model:selected-nodes="selectNodes"
    >
    </v-network-graph>
  </div>
  <div class="d-flex flex-row justify-end" style="width: 90%">
    <v-btn
      @click="testSend"
      color="purple-lighten-2"
      size="x-large"
      class="font-weight-bold"
      >送信</v-btn
    >
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
