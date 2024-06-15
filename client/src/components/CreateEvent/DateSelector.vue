<script setup lang="ts">
import { computed, ref, watch } from "vue";
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

watch([startTime, endTime], () => {
  console.log(startTime.value, endTime.value);
  if (startTime.value && endTime.value && startTime.value > endTime.value) {
    endTime.value = startTime.value;
  }
  console.log(startTime.value, endTime.value);
});

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
    <v-card class="h-100">
      <v-container
        class="d-flex justify-center height-100"
        :class="$style.container"
      >
        <div>
          <v-date-picker
            v-model="dates"
            label="日付"
            :allowedDates="allowedDates"
            multiple
            :class="$style.datePicker"
          />
          <div>
            <v-btn :class="$style.clear" @click="dates = []"
              >日付をクリア</v-btn
            >
          </div>
        </div>
        <div :class="$style.timeContainerWrap" class="w-50">
          <v-container
            class="d-flex flex-column justify-end py-0 px-0 mr-0 ml-4"
            :class="$style.timeContainer"
          >
            <v-list :class="$style.scroll" class="hidden-sm-and-down">
              <v-list-item
                v-for="date in sortedDates"
                :key="date.toISOString()"
                :title="format(date, 'yyyy/MM/dd')"
              />
            </v-list>
            <div :class="$style.timeRange">
              <v-text-field v-model="startTime" label="開始時刻" type="time" />
              <v-text-field v-model="endTime" label="終了時刻" type="time" />
            </div>
            <v-btn @click="updateValue(dates)">OK</v-btn>
          </v-container>
        </div>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<style module lang="scss">
.container {
  flex-direction: row;
  @media (max-width: 800px) {
    flex-direction: column;
  }
}
.timeRange {
  display: flex;
  align-items: center;
  justify-content: center;
}
.scroll {
  flex: 1;
  overflow-y: auto;
}
.timeContainerWrap {
  @media (max-width: 800px) {
    padding-top: 16px;
    padding-bottom: 16px;
    margin: auto;
  }
  @media (min-width: 800px) {
    margin-top: 0px;
    max-height: 100%;
    position: relative;
  }
}
.timeContainer {
  @media (min-width: 800px) {
    position: absolute;
    top: 0;
    bottom: 0;
  }
}

.datePicker {
  margin: auto;
}
.clear {
  display: block;
  margin: auto;
}
</style>
