<script setup lang="ts">
import { mdiHome } from "@mdi/js";
</script>

<template>
  <div>
    <v-app-bar color="primary">
      <v-app-title class="mx-6">hoge</v-app-title>
      <template v-slot:append>
        <v-btn @click="$router.push({ name: 'Home' })">
          <v-icon :icon="mdiHome" />
        </v-btn>
        <v-btn variant="text" @click="$router.push({ name: 'Events' })"
          >イベント一覧</v-btn
        >
        <v-btn variant="text" @click="$router.push({ name: 'CreateEvent' })"
          >イベント作成</v-btn
        >
      </template>
    </v-app-bar>
  </div>
</template>

<style module lang="scss"></style>
