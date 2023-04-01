import { createApp } from "vue";
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
axios.defaults.withCredentials = true;

const app = createApp(App);

app.use(router);
app.use(vuetify);

app.mount("#app");
