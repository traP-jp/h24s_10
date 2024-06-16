<script setup lang="ts">
import { ref } from "vue";
import { useGetEventsMe } from "/@/generated/api/openapi";
import { computed } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const { data: eventsMe } = useGetEventsMe();

const invitedEvents = computed(
  () =>
    eventsMe?.value?.data.filter(
      (v) => !v.isConfirmed && !v.isAnswered && !v.isHost
    ) ?? []
);

const hostEvents = computed(
  () => eventsMe?.value?.data.filter((v) => v.isHost) ?? []
);

const answeredEvents = computed(
  () => eventsMe?.value?.data.filter((v) => v.isAnswered) ?? []
);

const confirmedEvents = computed(
  () => eventsMe?.value?.data.filter((v) => v.isConfirmed) ?? []
);
</script>

<template>
  <div class="h-2">イベント一覧</div>
  <div class="eventList">
    <div>
      <div v-if="invitedEvents.length > 0">
        <h3 class="title">招待中</h3>
        <v-card
          v-for="event in invitedEvents"
          :key="event.event_id"
          class="ma-5"
          link
          @click="router.push(`/events/${event.event_id}/applicate-event`)"
        >
          <v-card-title>{{ event.title }}</v-card-title>
          <v-card-text>{{ event.description }}</v-card-text>
        </v-card>
      </div>
      <div v-if="hostEvents.length > 0">
        <h3 class="title">募集中</h3>
        <v-card
          v-for="event in hostEvents"
          :key="event.event_id"
          class="ma-5"
          link
          @click="router.push(`/events/${event.event_id}/applicate-event`)"
        >
          <v-card-title>{{ event.title }}</v-card-title>
          <v-card-text>{{ event.description }}</v-card-text>
        </v-card>
      </div>
      <div v-if="confirmedEvents.length > 0">
        <h3 class="title">参加イベント</h3>
        <v-card
          v-for="event in confirmedEvents"
          :key="event.event_id"
          class="ma-5"
          link
          @click="router.push(`/events/${event.event_id}`)"
        >
          <v-card-title>{{ event.title }}</v-card-title>
          <v-card-text>{{ event.description }}</v-card-text>
        </v-card>
      </div>
      <div v-if="answeredEvents.length > 0">
        <h3 class="title">回答済み</h3>
        <v-card
          v-for="event in answeredEvents"
          :key="event.event_id"
          class="ma-5"
          link
          @click="router.push(`/events/${event.event_id}`)"
        >
          <v-card-title>{{ event.title }}</v-card-title>
          <v-card-text>{{ event.description }}</v-card-text>
        </v-card>
      </div>
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
