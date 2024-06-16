<script setup lang="ts">
import { computed, ref } from "vue";
import type { VAutocomplete } from "vuetify/lib/components/index.mjs";
import DateSelector from "./DateSelector.vue";
import {
  mdiPlus,
  mdiClose,
  mdiAccountMultiple,
  mdiDotsHorizontal,
} from "@mdi/js";
import { format } from "date-fns";
import {
  usePostEvents,
  useGetMe,
  useGetTraqUsers,
  useGetTraqGroups,
} from "/@/generated/api/openapi";
import type { TraQUser, TraQGroup } from "/@/generated/api/openapi";
import { useRouter } from "vue-router";
import { title } from "process";

const router = useRouter();

const { isLoading, data: me } = useGetMe();
const { data: traQUsers } = useGetTraqUsers();
const { data: traQGroups } = useGetTraqGroups();
const { mutateAsync: postEvent } = usePostEvents();

type FilterFunction = Exclude<
  VAutocomplete["$props"]["customFilter"],
  undefined
>;
type Invitee = (TraQUser & { type: "user" }) | (TraQGroup & { type: "group" });

const users = computed<Invitee[]>(
  () => traQUsers.value?.data.map((v) => ({ ...v, type: "user" })) ?? []
);
const groups = computed<Invitee[]>(
  () => traQGroups.value?.data.map((v) => ({ ...v, type: "group" })) ?? []
);

const eventName = ref("");
const eventDescription = ref("");
const invitees = ref<Invitee[]>([]);

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

const snackbarTitle = ref(false);
const snackbarDates = ref(false);

const createEvent = async () => {
  if (eventName.value === "") {
    snackbarTitle.value = true;
    return;
  }
  if (Dates.value.length === 0) {
    snackbarDates.value = true;
    return;
  }
  await postEvent({
    data: {
      title: eventName.value,
      description: eventDescription.value,
      dateOptions: Dates.value.map(({ startDate, endDate }) => ({
        start: startDate.toISOString(),
        end: endDate.toISOString(),
      })),
      targets:
        invitees.value.length > 0
          ? invitees.value.flatMap((v) => {
              if (v.type === "user") {
                return [v.name];
              } else {
                return v.members?.map(({ name }) => name) ?? [];
              }
            })
          : [],
    },
  });
  console.log("created");
  router.push("/");
};
</script>

<template>
  <h1 class="text-h3 text-left">イベント作成</h1>
  {{ isLoading }}
  {{ me?.data.traQID }}
  <h2 class="text-h4 text-left">タイトル</h2>
  <v-text-field v-model="eventName" label="イベント名" />

  <h2 class="text-h4 text-left">概要</h2>
  <v-textarea v-model="eventDescription" label="イベント概要" />
  <h2 class="text-h4 text-left">招待者</h2>

  <v-autocomplete
    v-model="invitees"
    :items="[...users, ...groups]"
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
        v-if="item.raw.type === 'user'"
        v-bind="props"
        :prepend-avatar="`https://q.trap.jp/api/v3/public/icon/${item.raw.name}`"
        :text="item.raw.name"
      />
      <v-chip
        v-if="item.raw.type === 'group'"
        v-bind="props"
        :prepend-icon="mdiAccountMultiple"
        >{{ item.raw.name }} {{ item.raw.members?.length }}人</v-chip
      >
    </template>
    <template v-slot:item="{ props, item }">
      <v-list-item
        v-if="item.raw.type === 'user'"
        v-bind="props"
        :prepend-avatar="`https://q.trap.jp/api/v3/public/icon/${item.raw.name}`"
        :subtitle="item.raw.displayName"
        :title="item.raw.name"
      />
      <v-list-item
        v-if="item.raw.type === 'group'"
        v-bind="props"
        :prepend-icon="mdiAccountMultiple"
        :subtitle="item.raw.members?.map((v) => v.name).join(', ')"
        :title="item.raw.name"
      >
        <template v-slot:subtitle
          ><v-avatar
            v-for="(v, index) in item.raw.members?.slice(0, 6)"
            :key="v.name"
            size="32"
            color="blue"
            text-color="white"
          >
            <v-img
              :src="`https://q.trap.jp/api/v3/public/icon/${v.name}`"
              :alt="v.name"
            />
          </v-avatar>
          <v-icon
            v-if="item.raw.members?.length && item.raw.members?.length > 6"
            size="32"
            :icon="mdiDotsHorizontal"
          />
          {{ item.raw.members?.length }}人
        </template>
      </v-list-item>
    </template>
  </v-autocomplete>

  <h2 class="text-h4 text-left">日程候補</h2>
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
  <div class="pt-6"></div>
  <v-btn size="x-large" color="blue" @click="createEvent">イベント作成</v-btn>
  <div class="pt-12"></div>

  <v-snackbar v-model="snackbarTitle">
    タイトルが入力されていません。

    <template v-slot:actions>
      <v-btn
        color="pink"
        variant="text"
        @click="snackbarTitle = false"
        timeout="2000"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>

  <v-snackbar v-model="snackbarDates">
    日程が追加されていません。

    <template v-slot:actions>
      <v-btn
        color="pink"
        variant="text"
        @click="snackbarDates = false"
        timeout="2000"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
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
