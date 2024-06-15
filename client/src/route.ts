import { createRouter, createWebHashHistory } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import EventList from "./components/EventList.vue";
import Event from "./components/Event.vue";

export const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HelloWorld,
    },
    {
      path: "/events",
      name: "Events",
      component: EventList,
    },
    {
      path: "/events/:id",
      name: "EventDetail",
      component: Event,
      props: true,
    }
  ],
});
