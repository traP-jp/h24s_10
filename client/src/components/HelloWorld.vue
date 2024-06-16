<script setup lang="ts">
import { ref } from "vue";
import { mdiChevronRight } from "@mdi/js";
const count = ref(0);
</script>

<template>
  <v-card title="イベント一覧"
    class="btn"
    link
    @click="$router.push({name: 'Events'})"
    :append-icon="mdiChevronRight"
    subtitle="イベント一覧ページへ"/>
  <v-card title="イベント作成"
    class="btn"
    link
    @click="$router.push({name: 'CreateEvent'})"
    :append-icon="mdiChevronRight"
    subtitle="イベント作成ページへ"/>
</template>

<style scoped>
.read {
  color: #888;
}

.btn {
  margin: 10px;
  padding: 10px;
  text-align: left;
}
</style>
