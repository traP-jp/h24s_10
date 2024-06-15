<script setup lang="ts">
import Navigation from "./components/Navigation.vue";
</script>

<template>
  <div style="position: absolute; top: 0; left: 0; width: 100vw">
    <Navigation />
    <div style="max-width: 50rem; margin: auto; margin-top: 6rem">
      <RouterView />
    </div>
  </div>
</template>

<style module lang="scss"></style>
