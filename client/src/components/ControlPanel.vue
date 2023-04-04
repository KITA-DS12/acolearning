<script setup lang="ts">
import { useStore } from "@/store";
import { computed } from "vue";

const store = useStore();
const zoomLevel = computed({
  get: () => store.zoomLevel,
  set: (newValue) => store.setLevel(newValue),
});
const selectNodes = computed({
  get: () => store.selectedNodes,
  set: (newArray) => store.setSelectedNodes(newArray),
});

const addNode = () => {
  const nextNode = Object.keys(store.nodes).length;
  const nodeId: string = `node${nextNode}`;
  const name: string = `Node ${nextNode}`;
  store.addNode(nodeId, name);
};

const removeNode = () => {
  for (const nodeId of selectNodes.value) {
    store.removeNode(nodeId);
  }
};
</script>

<template>
  <div class="graph-space">
    <div class="control-panel d-flex flex-row justify-center">
      <label>Node:</label>
      <div class="button-block">
        <v-btn class="ma-1" variant="outlined" rounded="lg" @click="addNode">
          <v-icon color="deep-purple">mdi-plus-circle-outline</v-icon>
        </v-btn>
        <v-btn class="ma-1" variant="outlined" rounded="lg" @click="removeNode">
          <v-icon color="deep-purple">mdi-minus-circle-outline</v-icon>
        </v-btn>
      </div>
      <label>Layout:</label>
      <div class="button-block">
        <v-btn class="ma-1" variant="outlined" rounded="lg">
          <v-icon color="deep-purple">mdi-format-align-center</v-icon>
        </v-btn>
        <v-btn class="ma-1" variant="outlined" rounded="lg">
          <v-icon color="deep-purple">mdi-fullscreen</v-icon>
        </v-btn>
        <v-btn class="ma-1" variant="outlined" rounded="lg">
          <v-icon color="deep-purple">mdi-format-horizontal-align-right</v-icon>
        </v-btn>
        <v-btn class="ma-1" variant="outlined" rounded="lg">
          <v-icon color="deep-purple">mdi-format-vertical-align-bottom</v-icon>
        </v-btn>
      </div>
      <label>Zoom:</label>
      <div class="button-block">
        <v-slider v-model="zoomLevel" min="0.1" max="12" :step="1" color="deep-purple" track-color="deep-purple"
          style="width: 200px; margin-top: 5px; margin-bottom: -15px" />
      </div>
    </div>
  </div>
</template>
<style>
.button-block {
  width: fit-content;
  border: 2px solid;
  border-radius: 50px;
  border-color: #7e57c2;
  margin-top: 8px;
  padding: 10px 20px;
  text-align: center;
}

.control-panel>label {
  color: #4527a0;
  padding: 5px;
  padding-left: 15px;
  padding-top: 25px;
}
</style>
