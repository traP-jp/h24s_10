import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { router } from "./route";

import "vuetify/styles";
import { createVuetify } from "vuetify";

const vuetify = createVuetify();

createApp(App).use(router).use(vuetify).mount("#app");
