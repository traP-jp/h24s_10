<script setup lang="ts">

import { ref } from 'vue';

defineProps(['id']);

interface Event {
    title: string
    description: string
    date: string[]
    participant: string[][]
    id: number
}

const event = ref<Event>(
    {
        title: 'hoge',
        description: 'fuga',
        date: [
            '2024/06/17 16:00-18:00',
            '2024/06/18 16:00-18:00'
        ],
        participant:[
            [
                "pirosiki",
                "noc7t",
                "jippo"
            ],
            [
                "pippi0057",
                "ogu_kazemiya",
                "Luftalian"
            ]
        ],
        id: 1
    }
)

const imgPath = (name: string) => {
    return `https://q.trap.jp/api/v3/public/icon/${name}`
}

</script>

<template>
    <div class="event">
        <div class="title">{{ event.title }}</div>
        <v-divider :thickness="2"/>
        <v-container class="date w-100">
            <div class="text-h6">日時・参加者</div>
            <v-divider />
            <div v-for="d, idx in event.date" :key="d" class="text-h6">
                <p>{{ d }}</p>
                <v-avatar v-for="p in event.participant[idx]" :key="p" class="avatar">
                    <v-img :src="imgPath(p)"></v-img>
                </v-avatar>
            </div>
        </v-container>
        <v-container class="detail">
            <div class="text-h6">概要</div>
            <v-divider />
            <div class="description">{{ event.description }}</div>
        </v-container>
    </div>
    <v-btn color="green" @click="$router.go(-1)">戻る</v-btn>
</template>

<style scoped>
.title {
    font-size: 30px;
    margin: 15px;
}

.date {
    text-align: left;
    margin: 10px;
    justify-content: left;

}

.detail {
    text-align: left;
    margin: 10px;
    justify-content: left;
}

.participant {
    text-align: left;
    margin: 10px;
    justify-content: right;
}

.description{
    font-size: 20px;
}

.avatar {
    margin: 3px;
}
</style>