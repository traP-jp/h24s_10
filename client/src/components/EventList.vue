<script setup lang="ts">
import { ref } from "vue";

interface Event {
  title: string;
  description: string;
  date: string;
  id: number;
}

const events = ref<Event[]>([
  {
    title: "hoge",
    description: "fuga",
    date: "2024/06/17",
    id: 1,
  },
]);

const participatedEvents = ref<Event[]>([
  {
    title: "hoge",
    description: "fuga",
    date: "2024/06/15",
    id: 1,
  },
  {
    title: "foo",
    description: "bar",
    date: "2024/06/16",
    id: 2,
  },
]);
</script>

<template>
  <div class="head">イベント一覧</div>
  <div class="eventList">
    <div>
      <h3 class="title">参加予定</h3>
      <v-card
        v-for="event in participatedEvents"
        :key="event.id"
        class="ma-5"
        link
        @click="$router.push({ name: 'EventDetail', params: { id: event.id } })"
      >
        <v-card-title>{{ event.title }}</v-card-title>
        <v-card-subtitle>{{ event.date }}</v-card-subtitle>
        <v-card-text>{{ event.description }}</v-card-text>
      </v-card>
    </div>
    <div>
      <h3 class="title">募集中</h3>
      <v-card v-for="event in events" :key="event.title" class="ma-5" link>
        <v-card-title>{{ event.title }}</v-card-title>
        <v-card-subtitle>{{ event.date }}</v-card-subtitle>
        <v-card-text>{{ event.description }}</v-card-text>
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
