<script setup lang="ts">
import { computed, ref, watch } from "vue";
import {
  format,
  startOfDay,
  setMinutes,
  setHours,
  setSeconds,
  setMilliseconds,
} from "date-fns";

const emit = defineEmits<{
  (e: "update:value", value: { startDate: Date; endDate: Date }[]): void;
  (e: "close"): void;
}>();

const now = new Date();
const dates = ref<Date[]>([]);
const sortedDates = computed(() =>
  dates.value.sort((a, b) => a.getTime() - b.getTime())
);

const startTime = ref(format(now, "HH:mm"));
const endTime = ref(format(now, "HH:mm"));

watch([startTime, endTime], () => {
  if (startTime.value && endTime.value && startTime.value > endTime.value) {
    endTime.value = startTime.value;
  }
});

const onClickOK = (value: Date[]) => {
  console.log(value, startTime.value, endTime.value);
  const [startHour, startMinute] = startTime.value.split(":");
  const [endHour, endMinute] = endTime.value.split(":");
  const selectedDates = sortedDates.value.map((date) => {
    date = setSeconds(setMilliseconds(date, 0), 0);
    return {
      startDate: setHours(
        setMinutes(date, Number(startMinute)),
        Number(startHour)
      ),
      endDate: setHours(setMinutes(date, Number(endMinute)), Number(endHour)),
    };
  });
  console.log(selectedDates);
  emit("update:value", selectedDates);
  emit("close");
};
const onClickClose = () => {
  emit("close");
};

const allowedDates = (value: any) => {
  const today = startOfDay(now);
  return value >= today;
};
</script>

<template>
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
          <v-btn :class="$style.clear" @click="dates = []">日付をクリア</v-btn>
        </div>
      </div>
      <div :class="$style.timeContainerWrap" class="w-50">
        <v-container
          class="d-flex flex-column justify-end py-0 px-0 mr-0"
          :class="$style.timeContainer"
        >
          <v-list :class="$style.scroll" class="hidden-sm-and-down">
            <TransitionGroup name="list">
              <v-list-item
                v-for="date in sortedDates"
                :key="date.toISOString()"
                :title="format(date, 'yyyy/MM/dd')"
              />
            </TransitionGroup>
          </v-list>
          <div :class="$style.timeRange">
            <v-text-field v-model="startTime" label="開始時刻" type="time" />
            <v-text-field v-model="endTime" label="終了時刻" type="time" />
          </div>
          <v-row class="justify-space-around" :class="$style.buttonRow">
            <v-col>
              <v-btn block @click="onClickClose()">閉じる</v-btn>
            </v-col>
            <v-col>
              <v-btn block @click="onClickOK(dates)">OK</v-btn>
            </v-col>
          </v-row>
        </v-container>
      </div>
    </v-container>
  </v-card>
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
    margin-left: 16px;
    position: absolute;
    top: 0;
    bottom: 0;
  }
}

.buttonRow {
  display: flex;
  flex-grow: 0 !important;
}

.datePicker {
  margin: auto;
}
.clear {
  display: block;
  margin: auto;
}
</style>

<style scoped>
.list-move, /* 移動する要素にトランジションを適用 */
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>
