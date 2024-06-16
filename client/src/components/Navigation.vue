<script setup lang="ts">
import { mdiHome } from "@mdi/js";
import { useRouter } from "vue-router";

const router = useRouter();
</script>

<template>
  <div>
    <v-app-bar color="primary">
      <v-app-bar-title class="mx-6">予定決めアプリ</v-app-bar-title>
      <template v-slot:append>
        <v-btn @click="router.push('/')">
          <v-icon :icon="mdiHome" />
        </v-btn>
        <v-btn variant="text" @click="router.push('/events')"
          >イベント一覧</v-btn
        >
        <v-btn variant="text" @click="router.push('/create-event')"
          >イベント作成</v-btn
        >
      </template>
    </v-app-bar>
  </div>
</template>

<style module lang="scss"></style>
