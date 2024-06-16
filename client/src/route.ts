import { createRouter, createWebHistory } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import Event from "./components/Event.vue";
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
      path: "/events",
      name: "Events",
      component: () => import("./components/EventMyList.vue"),
    },
    {
      path: "/events/all",
      name: "AllEvents",
      component: () => import("./components/EventAllList.vue"),
    },

    {
      path: "/events/:id/detail",
      name: "EventDetail",
      component: () => import("./components/EventDetail.vue"),
    },
    {
      path: "/events/:id",
      name: "Event",
      component: Event,
    },
    {
      path: "/create-event",
      name: "CreateEvent",
      component: () => import("./components/CreateEvent/CreateEvent.vue"),
    },
  ],
});
