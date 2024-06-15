<script setup lang="ts">
import{ ref } from 'vue'

const eventTitle = ref('イベントタイトル')
const description = ref('説明をここに書く イベントの説明とかいろいろ')
const organizer = ref('主催者')

</script>

<template>
  <div>
    <h1>イベント詳細</h1>
    <div>
      <h2>{{ eventTitle }}</h2>
      <div>by {{ organizer }}</div>
      <div>{{ description }}</div>
      <div>日程選択するところ(未実装)</div>
      <div>
        参加予定の人一覧(未実装)
      </div>
    </div>
  </div>
</template>

<style>
</style>
