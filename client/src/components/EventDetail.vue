<script setup lang="ts">
import{ ref } from "vue";
import { format, parseISO, startOfDay } from "date-fns";
import ApplicateEvent from "./ApplicateEvent.vue";
import DecideDate from "./DecideDate.vue";

type DateOption = {
  dateOptionID: string;
  start: Date;
  end: Date;
}
type EventDetail = {
  traqID: string;
  title: string;
  organizer: string;
  description: string;
  isConfirmed: boolean;
  dateOptions: DateOption[];
}

// exsample event
const eventDetail: EventDetail = {
  traqID: "eventID",
  title: "eventTitle",
  organizer: "ogu_kazemiya",
  description: "eventDescription",
  isConfirmed: false,
  dateOptions: [
    { dateOptionID: "dateOptionID0", start: new Date(), end: new Date() },
    { dateOptionID: "dateOptionID1", start: new Date(), end: new Date() },
  ],
}

const isOrganizer = ref<boolean>(true)
const isDecidingMode = ref<boolean>(true)
</script>

<template>
  <div>
    <h1>イベント詳細</h1>
    <div>
      <h2>{{ eventDetail.title }}</h2>
      <div>
        by {{ eventDetail.organizer }}
        <img 
          :src="`https://q.trap.jp/api/v3/public/icon/${eventDetail.organizer}`"
          :alt="`${eventDetail.organizer}'s icon'`"
          height="25" width="25" />
      </div>
      <div>{{ eventDetail.description }}</div>
      <ApplicateEvent v-if="!isDecidingMode" :eventDetail="eventDetail" />
      <DecideDate v-if="isDecidingMode" :eventDetail="eventDetail" />
    </div>
  </div>
</template>

<style>
</style>
