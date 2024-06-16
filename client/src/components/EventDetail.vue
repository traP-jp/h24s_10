<script setup lang="ts">
import { computed, ref } from "vue";
import ApplicateEvent from "./ApplicateEvent.vue";
import DecideDate from "./DecideDate.vue";
import Event from "./Event.vue";

import { useRoute } from "vue-router";
import { watch } from "vue";
import { useGetEventsEventID, useGetMe } from "../generated/api/openapi";

const route = useRoute();

const id = Array.isArray(route.params.id)
  ? route.params.id[0]
  : route.params.id;

const { data: eventsAxios } = useGetEventsEventID(id);
const { data: me } = useGetMe();

const event = computed(() => eventsAxios.value?.data);

const dateOptionIDs = ref<boolean[]>([]);

watch(event, () => {
  dateOptionIDs.value = event.value?.dateOptions?.map((v) => false) ?? [];
});

const isHost = computed<boolean>(
  () => event.value?.hostID === me.value?.data.traQID
);
</script>

<template>
  <div>
    <h1>{{ event?.title }}</h1>
    <div>
      <div>
        by {{ event?.hostID }}
        <img
          :src="`https://q.trap.jp/api/v3/public/icon/${event?.hostID}`"
          :alt="`${event?.hostID}'s icon'`"
          height="25"
          width="25"
        />
      </div>
      <div>{{ event?.description }}</div>
      <DecideDate v-if="isHost && !event?.isConfirmed" />
      <ApplicateEvent v-else-if="!event?.isConfirmed" />
      <Event v-else />
    </div>
  </div>
</template>

<style></style>
