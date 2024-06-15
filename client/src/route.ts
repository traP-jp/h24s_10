import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import ApplicateEvent from "./components/ApplicateEvent.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HelloWorld,
    },
    {
      path: "/ApplicateEvent",
      name: "ApplicateEvent",
      component: ApplicateEvent,
    },
    {
      path: "/create-event",
      name: "CreateEvent",
      component: () => import("./components/CreateEvent/CreateEvent.vue"),
    },
  ],
});
