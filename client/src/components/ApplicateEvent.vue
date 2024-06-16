<script setup lang="ts">
import{ ref } from 'vue';
import { format, parseISO, startOfDay } from "date-fns";

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

defineProps<{ eventDetail: EventDetail }>()

const dateOptionIDs = ref<Date[]>([])

const postDateOptionIDs = () => {
  // DateOptionIDsをPOSTする
}
</script>

<template>
  <div v-if="!eventDetail.isConfirmed">
    <div v-for="dateOption in eventDetail.dateOptions" :key="dateOption.dateOptionID">
      <v-checkbox
        v-model="dateOptionIDs"
        :value="dateOption.dateOptionID"
        :label="`${format(dateOption.start, 'MM/dd(eee)')}
          ${format(dateOption.start, 'HH:mm')} ~ ${format(dateOption.end, 'HH:mm')}`"
        hide-details>
      </v-checkbox>
    </div>
    <v-btn @click="postDateOptionIDs" color="green">保存</v-btn>
  </div>

  <div v-if="eventDetail.isConfirmed">
    <div>
      {{ format(eventDetail.dateOptions[0].start, 'MM/dd(eee)') }}
      {{ format(eventDetail.dateOptions[0].start, 'HH:mm') }} ~
      {{ format(eventDetail.dateOptions[0].end, 'HH:mm') }}
    </div>
  </div>
</template>

<style>
</style>
