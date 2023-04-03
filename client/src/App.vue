<script setup lang="ts">
import Nav from "/src/components/Nav.vue";
import ControlPanel from "/src/components/ControlPanel.vue";
import GraphSpace from "/src/components/GraphSpace.vue";
</script>

<template>
  <v-app>
    <Nav />
    <div style="height: 200px" />
    <ControlPanel />
    <div style="height: 20px" />
    <GraphSpace />
  </v-app>
</template>

<style scoped></style>
