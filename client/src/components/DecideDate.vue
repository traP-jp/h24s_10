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
const props = defineProps<{eventDetail: EventDetail}>()

const dateID = ref()
const applicantDateList = ref<{ traqID: string, dateOptionIDs: string[] }[]>([
  {traqID:"ogu_kazemiya", dateOptionIDs:["dateOptionID0"]}
])
const applicantsMap = ref(new Map<string, string[]>()) // dateOptionID -> applicantsID のmapping

// はじめに実行
const calcDateOptionApplicants = () => {
  // applicantDateListを/events/{eventID}/applicantsで取得
  for (const dateOption of props.eventDetail.dateOptions) {
    applicantsMap.value.set(dateOption.dateOptionID, [])
  }
  for (const applicant of applicantDateList.value) {
    for (const dateOptionID of applicant.dateOptionIDs) {
      applicantsMap.value.get(dateOptionID)?.push(applicant.traqID)
    }
  }
}
</script>

<template>
  <div>
    <div v-if="!eventDetail.isConfirmed">
      <v-radio-group v-model="dateID">
        <div v-for="dateOption in eventDetail.dateOptions" :key="dateOption.dateOptionID" class="selectDateOption">
          <img v-for="applicantID in applicantsMap.get(dateOption.dateOptionID)" :key="applicantID"
            :src="`https://q.trap.jp/api/v3/public/icon/${applicantID}`"
            :alt="`${applicantID}'s icon'`"
            height="25" width="25" />
          <v-radio 
            :value="dateOption.dateOptionID"
            :label="`${format(dateOption.start, 'MM/dd(eee)')}
              ${format(dateOption.start, 'HH:mm')} ~ ${format(dateOption.end, 'HH:mm')}`"
            hide-details
          >
          </v-radio>
        </div>
      </v-radio-group>
    </div>

    <div v-if="eventDetail.isConfirmed">
      <div>
        {{ format(eventDetail.dateOptions[0].start, 'MM/dd(eee)') }}
        {{ format(eventDetail.dateOptions[0].start, 'HH:mm') }} ~
        {{ format(eventDetail.dateOptions[0].end, 'HH:mm') }}
      </div>
    </div>
  </div>
</template>

<style>
.selectDateOption{
  display: flex;
}
</style>
