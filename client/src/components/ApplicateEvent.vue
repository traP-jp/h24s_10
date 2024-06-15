<script setup lang="ts">
import { ref } from "vue";
import { format, parseISO, startOfDay } from "date-fns";

type DateOption = {
  id: string;
  start: Date;
  end: Date;
};
type EventDetail = {
  id: string;
  title: string;
  organizer: string;
  description: string;
  isConfirmed: boolean;
  dateOptions: DateOption[];
};

// exsample event
const eventDetail: EventDetail = {
  id: "eventID",
  title: "eventTitle",
  organizer: "ogu_kazemiya",
  description: "eventDescription",
  isConfirmed: false,
  dateOptions: [{ id: "dateOptionID", start: new Date(), end: new Date() }],
};

const dateOptionIDs = ref<Date[]>([]);

const postDateOptionIDs = () => {
  // DateOptionIDsをPOSTする
};
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
          height="25"
          width="25"
        />
      </div>
      <div>{{ eventDetail.description }}</div>

      <div v-if="!eventDetail.isConfirmed">
        <div v-for="dateOption in eventDetail.dateOptions" :key="dateOption.id">
          <v-checkbox
            v-model="dateOptionIDs"
            :value="dateOption.id"
            :label="`${format(dateOption.start, 'MM/dd(eee)')}
              ${format(dateOption.start, 'HH:mm')} ~ ${format(
              dateOption.end,
              'HH:mm'
            )}`"
            hide-details
          >
          </v-checkbox>
        </div>
        <v-btn @click="postDateOptionIDs" color="green">保存</v-btn>
      </div>

      <div v-if="eventDetail.isConfirmed">
        <div>
          {{ format(eventDetail.dateOptions[0].start, "MM/dd(eee)") }}
          {{ format(eventDetail.dateOptions[0].start, "HH:mm") }} ~
          {{ format(eventDetail.dateOptions[0].end, "HH:mm") }}
        </div>
      </div>
    </div>
  </div>
</template>

<style></style>
