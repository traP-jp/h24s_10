<script setup lang="ts">
import { computed, ref } from "vue";
import type { VAutocomplete } from "vuetify/lib/components/index.mjs";
import DateSelector from "./DateSelector.vue";
import { mdiPlus, mdiClose } from "@mdi/js";
import { format } from "date-fns";
import {
  usePostEvents,
  useGetMe,
  useGetTraqUsers,
  useGetTraqGroups,
} from "/@/generated/api/openapi";
import { start } from "repl";
import { onMounted } from "vue";

const { isLoading, data: me } = useGetMe();
const { data: traQusers } = useGetTraqUsers();
const { data: groups } = useGetTraqGroups();
const { mutate: postEvent, error: postEventError } = usePostEvents();

type FilterFunction = Exclude<
  VAutocomplete["$props"]["customFilter"],
  undefined
>;
type User = {
  name: string;
  displayName: string;
};

const users = computed(() => traQusers.value?.data ?? []);

const eventName = ref("");
const eventDescription = ref("");
const invitees = ref<User[]>([]);

const userSearchFilter: FilterFunction = (_, query, item?) => {
  const name = item?.raw.name ?? "";
  const displayName = item?.raw.displayName ?? "";

  return (
    name.toLowerCase().includes(query.toLowerCase()) ||
    displayName.toLowerCase().includes(query.toLowerCase())
  );
};

const dialog = ref(false);
const Dates = ref<{ startDate: Date; endDate: Date }[]>([]);

const addDates = (dates: { startDate: Date; endDate: Date }[]) => {
  Dates.value = [...Dates.value, ...dates];
  Dates.value.sort((a, b) => a.startDate.getTime() - b.startDate.getTime());
  console.log(Dates.value[0].startDate.toISOString());
};

const removeDate = (index: number) => {
  Dates.value.splice(index, 1);
};

const createEvent = () => {
  console.log({
    title: eventName.value,
    description: eventDescription.value,
    dateOptions: Dates.value.map(({ startDate, endDate }) => ({
      start: startDate.toISOString(),
      end: endDate.toISOString(),
    })),
    targets: invitees.value.map(({ name }) => name),
  });
  postEvent({
    data: {
      title: eventName.value,
      description: eventDescription.value,
      dateOptions: Dates.value.map(({ startDate, endDate }) => ({
        start: startDate.toISOString(),
        end: endDate.toISOString(),
      })),
      targets: invitees.value.map(({ name }) => name),
    },
  });
};
</script>

<template>
  <h1>イベント作成</h1>
  {{ isLoading }}
  {{ me?.data.traQID }}
  <h2>イベント名</h2>
  <v-text-field v-model="eventName" label="イベント名" />

  <h2>イベント概要</h2>
  <v-textarea v-model="eventDescription" label="イベント概要" />
  <h2>招待者</h2>

  <v-autocomplete
    v-model="invitees"
    :items="users"
    label="招待者"
    item-title="name"
    item-value="name"
    multiple
    auto-select-first
    closable-chips
    return-object
    :custom-filter="userSearchFilter"
  >
    <template v-slot:chip="{ props, item }">
      <v-chip
        v-bind="props"
        :prepend-avatar="`https://q.trap.jp/api/v3/public/icon/${item.raw.name}`"
        :text="item.raw.name"
      ></v-chip>
    </template>
    <template v-slot:item="{ props, item }">
      <v-list-item
        v-bind="props"
        :prepend-avatar="`https://q.trap.jp/api/v3/public/icon/${item.raw.name}`"
        :subtitle="item.raw.displayName"
        :title="item.raw.name"
      ></v-list-item>
    </template>
  </v-autocomplete>

  <h2>日程候補</h2>
  <v-list :class="$style.list">
    <TransitionGroup name="list">
      <v-list-item
        v-for="(date, index) in Dates"
        :key="date.startDate.toISOString() + date.endDate.toISOString()"
      >
        {{ format(date.startDate, "yyyy年MM月dd日") }}
        {{ format(date.startDate, "HH:mm") }}
        ～
        {{ format(date.endDate, "HH:mm") }}
        <v-btn @click="removeDate(index)" :icon="mdiClose" variant="plain" />
      </v-list-item>
    </TransitionGroup>
  </v-list>

  <v-btn @click="dialog = true">
    <v-icon :icon="mdiPlus" />
    日程を追加
  </v-btn>
  <v-dialog v-model="dialog" max-width="800" max-height="600">
    <DateSelector @close="dialog = false" @update:value="addDates" />
  </v-dialog>
  <br />
  <v-btn color="blue" @click="createEvent">イベント作成</v-btn>
  {{ postEventError }}
</template>

<style module lang="scss">
.list {
  position: relative;
  overflow: hidden;
  width: fit-content;
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
