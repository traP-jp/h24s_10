<script setup lang="ts">
import { ref } from 'vue';
import { mdiChevronRight } from "@mdi/js";

interface Event {
    title: string
    description: string
    date: string
    id: number
}

const events = ref<Event[]>([
    {
        title: 'hoge',
        description: 'fuga',
        date: '2024/06/17',
        id: 1
    }
])

const participatedEvents = ref<Event[]>([
    {
        title: 'hoge',
        description: 'fuga',
        date: '2024/06/15',
        id: 1
    },
    {
        title: 'foo',
        description: 'bar',
        date: '2024/06/16',
        id: 2
    }
])

const stringSlice = (str: string, length: number) => {
    if(str.length <= length) return str
    return str.slice(0, Math.min(str.length, length)) + '...'
}

</script>

<template>
    <div class="head">イベント一覧</div>
    <div class="eventList">
        <div>
            <h3 class="title">参加予定</h3>
            <v-card v-for="event in participatedEvents"
                :key="event.id"
                class="ma-5"
                link
                @click="$router.push({ name: 'EventDetail', params: { id: event.id } })"
                :append-icon="mdiChevronRight"
                :title="event.title"
                :subtitle="event.date">
                <v-card-text>{{ stringSlice(event.description, 50) }}</v-card-text>
            </v-card>
        </div>
        <div>
            <h3 class="title">募集中</h3>
            <v-card v-for="event in events"
                :key="event.title"
                class="ma-5"
                link
                @click="$router.push({ name: 'ApplicateEvent'})"
                :append-icon="mdiChevronRight"
                :title="event.title"
                :subtitle="event.date">
                <v-card-text>{{ stringSlice(event.description, 50) }}</v-card-text>
            </v-card>
        </div>
    </div>
</template>

<style scoped>
.head {
    font-size: 30px;
    margin: 10px;
}

.eventList {
    text-align: left;
}

.title {
    margin: 10px;
    margin-left: 25px;
}
</style>