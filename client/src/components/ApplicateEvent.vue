<script setup lang="ts">
import{ ref } from 'vue'

type DateOption = string

type EventDetail = {
  eventName: string;
  organizer: string;
  description: string;
  dateOptions: DateOption[];
}
type EntryStatus = { dateOption: DateOption, status: boolean }

const eventDetail: EventDetail = {
  eventName: "イベント",
  organizer: "ogu_kazemiya",
  description: "イベントの説明文",
  dateOptions: ["1/1 10:00-", "1/1 11:00-", "1/2 10:00-"],
}

// const entryStatus: EntryStatus =

const checkbox = ref<DateOption[]>([])
</script>

<template>
  <div>
    <h1>イベント詳細</h1>
    <div>
      <h2>{{ eventDetail.eventName }}</h2>
      <div>
        by {{ eventDetail.organizer }}
        <img 
          :src="`https://q.trap.jp/api/v3/public/icon/${eventDetail.organizer}`"
          :alt="`${eventDetail.organizer}'s icon'`"
          height="50" width="50" />
      </div>
      <div>{{ eventDetail.description }}</div>
      <div v-for="dateOption in eventDetail.dateOptions" :key="dateOption">
        <v-checkbox
          v-model="checkbox"
          :value="dateOption"
          :label="dateOption"
          hide-details>
        </v-checkbox>
      </div>
      <div>
        {{ checkbox }}
      </div>
    </div>
  </div>
</template>

<style>
</style>
