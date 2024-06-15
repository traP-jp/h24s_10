<script setup lang="ts">
import { computed, ref } from "vue";
import { format, parseISO, startOfDay } from "date-fns";

const emit = defineEmits<{
  (e: "update:value", value: Date[]): void;
}>();

const now = new Date();
console.log(now);
const dates = ref<Date[]>([now]);
const sortedDates = computed(() =>
  dates.value.toSorted((a, b) => a.getTime() - b.getTime())
);

const startTime = ref(format(now, "HH:mm"));
const endTime = ref(format(now, "HH:mm"));

const updateValue = (value: Date[]) => {
  console.log(value, startTime.value, endTime.value);
};

const allowedDates = (value: any) => {
  const today = startOfDay(now);
  return value >= today;
};
const dialog = ref(false);
</script>

<template>
  <v-btn @click="dialog = !dialog"></v-btn>
  <v-dialog v-model="dialog" max-width="800" max-height="600">
    <v-card>
      <v-container class="d-flex flex-row justify-center">
        <div>
          <v-date-picker
            v-model="dates"
            label="日付"
            :allowedDates="allowedDates"
            multiple
          />
          <v-btn block @click="dates = []">日付をクリア</v-btn>
        </div>
        <v-container
          class="w-50 d-flex flex-column justify-end py-0 px-0 mr-0 ml-4"
          :class="$style.timeContainer"
        >
          <div :class="$style.scroll">
            <div v-for="date in sortedDates" :key="date.toISOString()">
              {{ format(date, "yyyy/MM/dd") }}
            </div>
          </div>
          <div :class="$style.timeRange">
            <v-text-field
              v-model="startTime"
              label="開始時刻"
              type="time"
              :min="new Date().toISOString()"
            />
            <v-text-field
              v-model="endTime"
              label="終了時刻"
              type="time"
              :min="new Date().toISOString()"
            />
          </div>
          <v-btn @click="updateValue(dates)">OK</v-btn>
        </v-container>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<style module lang="scss">
.timeRange {
  display: flex;
  align-items: center;
  justify-content: center;
}
.scroll {
  flex: 1;
  overflow-y: auto;
}
.timeContainer {
  max-height: 100%;
}
</style>
