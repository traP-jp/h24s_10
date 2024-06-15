<script setup lang="ts">
import { ref } from "vue";
import type { VAutocomplete } from "vuetify/lib/components/index.mjs";
import DateSelector from "./DateSelector.vue";

type FilterFunction = Exclude<
  VAutocomplete["$props"]["customFilter"],
  undefined
>;
type User = {
  name: string;
  displayName: string;
};

const users: User[] = [
  { name: "noc7t", displayName: "ノクナートン" },
  { name: "ogu_kazemiya", displayName: "かぜみやおぐ" },
  { name: "pirosiki", displayName: "pirosiki" },
  { name: "Luftalian", displayName: "ルフタリアン" },
  { name: "jippo", displayName: "jippo" },
  { name: "pippi0057", displayName: "pippi0057" },
];

const count = ref(0);
const eventName = ref("");
const eventDescription = ref("");
const invitees = ref<User[]>([]);
const startDate = ref(new Date());

const userSearchFilter: FilterFunction = (_, query, item?) => {
  const name = item?.raw.name ?? "";
  const displayName = item?.raw.displayName ?? "";

  return (
    name.toLowerCase().includes(query.toLowerCase()) ||
    displayName.toLowerCase().includes(query.toLowerCase())
  );
};
</script>

<template>
  <h1>イベント作成</h1>
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

  <DateSelector />

  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
      >create-vue</a
    >, the official Vue + Vite starter
  </p>
  <p>
    Install
    <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p :class="$style.read">Click on the Vite and Vue logos to learn more</p>
  <v-btn @click="count++">count is {{ count }}</v-btn>
</template>

<style module lang="scss">
.read {
  color: #888;
}
</style>
