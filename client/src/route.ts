import { createRouter, createWebHashHistory } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";

export const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HelloWorld,
    },
    {
      path: "/create-event",
      name: "CreateEvent",
      component: () => import("./components/CreateEvent/CreateEvent.vue"),
    },
  ],
});
