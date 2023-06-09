import { createApp } from "vue";
import { createPinia } from "pinia";
import { useStore } from "./store";
import App from "./App.vue";
import router from "./router";

import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import "@mdi/font/css/materialdesignicons.css";

const vuetify = createVuetify({
  components,
  directives,
});

import axios from "axios";

axios.defaults.baseURL = "http://localhost:8888/";

import VNetworkGraph from "v-network-graph";
import "v-network-graph/lib/style.css";

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.provide("graph", useStore);
app.use(router);
app.use(vuetify);
app.use(VNetworkGraph);

app.mount("#app");
