<script setup lang="ts">
import { useGetEventsAll } from "/@/generated/api/openapi";
import { computed } from "vue";
import { useRouter } from "vue-router";
import { mdiChevronRight } from "@mdi/js";

const router = useRouter();

const { data: eventsAll } = useGetEventsAll();

const events = computed(() => eventsAll.value?.data);
</script>

<template>
  <div class="h-2">イベント一覧</div>
  <v-btn @click="router.push('/events')">自分に関係するイベント</v-btn>
  <div class="eventList">
    <div>
      <h3 class="title">全てのイベント</h3>
      <v-card
        v-for="event in events"
        :key="event.id"
        class="ma-5"
        link
        @click="$router.push(`/events/${event.id}/detail`)"
        :append-icon="mdiChevronRight"
        :title="event.title"
      >
      </v-card>
    </div>
  </div>
</template>

<style scoped>
.head {
  font-size: 30px;
  margin: 10px;
}

.eventList {
  text-align: left;
}

.title {
  margin: 10px;
  margin-left: 25px;
}
</style>
