<script setup lang="ts">
import Navigation from "./components/Navigation.vue";
</script>

<template>
  <v-app>
    <Navigation />
  </v-app>
  <div style="position: absolute; top: 0; left: 0; width: 100vw">
    <div style="max-width: 50rem; margin: auto; margin-top: 6rem">

      <RouterView />
    </div>
  </div>
</template>

<style module lang="scss"></style>
