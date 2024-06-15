import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { router } from "./route";

import { VueQueryPlugin } from "@tanstack/vue-query";

import "vuetify/styles";
import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi-svg";

const vuetify = createVuetify({
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
});

createApp(App).use(router).use(vuetify).use(VueQueryPlugin).mount("#app");
