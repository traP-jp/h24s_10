<script setup lang="ts">
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { format } from "date-fns";

import { useGetEventsEventID } from "/@/generated/api/openapi";
import { useGetEventsEventIDParticipants } from "/@/generated/api/openapi";
import { useGetEventsEventIDApplicants } from "/@/generated/api/openapi";

const route = useRoute();
const router = useRouter();

const id = Array.isArray(route.params.id)
  ? route.params.id[0]
  : route.params.id;

const { data: eventData } = useGetEventsEventID(id);
const { data: participants } = useGetEventsEventIDParticipants(id);
const { data: applicants } = useGetEventsEventIDApplicants(id);

const event = computed(() => eventData.value?.data);

const imgPath = (name: string) => {
  return `https://q.trap.jp/api/v3/public/icon/${name}`;
};
</script>

<template>
  <div class="event">
    <div class="text-h3 text-left">{{ event?.title }}</div>
    <v-divider :thickness="2" />
    <v-container class="date w-100">
      <div v-if="event?.date">
        <h2 class="text-h4">日時・参加者</h2>
        <v-divider />
        <div class="text-h6">
          {{ format(new Date(event?.date.start ?? ""), "yyyy年MM月dd日") }}
          {{ format(new Date(event?.date.start ?? ""), "HH:mm") }}
          ～
          {{ format(new Date(event?.date.end ?? ""), "HH:mm") }}
          <v-avatar v-for="p in participants?.data" :key="p" class="avatar">
            <v-img :src="imgPath(p)"></v-img>
          </v-avatar>
        </div>
      </div>
      <div v-else>
        <h2 class="text-h4">候補日程・投票者</h2>
        <v-divider />
        <div v-for="d in event?.dateOptions" :key="d.id" class="text-h6">
          {{ format(new Date(d.start ?? ""), "yyyy年MM月dd日") }}
          {{ format(new Date(d.start ?? ""), "HH:mm") }}
          ～
          {{ format(new Date(d.end ?? ""), "HH:mm") }}
          <v-avatar
            v-for="p in applicants?.data.filter((v) =>
              v.dateOptionIDs?.includes(d.id)
            )"
            :key="p.traqID"
            class="avatar"
          >
            <v-img :src="imgPath(p?.traqID ?? '')"></v-img>
          </v-avatar>
        </div>
      </div>
    </v-container>
    <v-container class="detail">
      <div class="text-h4">概要</div>
      <v-divider />
      {{ event?.description }}
    </v-container>
  </div>
  <v-btn color="green" @click="router.back()">戻る</v-btn>
</template>

<style scoped>
.title {
  font-size: 30px;
  margin: 15px;
}

.date {
  text-align: left;
  margin: 10px;
  justify-content: left;
}

.detail {
  text-align: left;
  margin: 10px;
  justify-content: left;
}

.participant {
  text-align: left;
  margin: 10px;
  justify-content: right;
}

.description {
  font-size: 20px;
}

.avatar {
  margin: 3px;
}
</style>
