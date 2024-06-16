<script setup lang="ts">
import { computed, ref } from "vue";
import { format } from "date-fns";

import {
  useGetEventsEventID,
  usePostEventsEventIDApplicants,
} from "../generated/api/openapi";

import { useRouter, useRoute } from "vue-router";
import { watch } from "vue";

const router = useRouter();
const route = useRoute();

const id = Array.isArray(route.params.id)
  ? route.params.id[0]
  : route.params.id;

const { data: eventsAxios } = useGetEventsEventID(id);
const { mutateAsync: postApplicants } = usePostEventsEventIDApplicants();

const event = computed(() => eventsAxios.value?.data);

const dateOptionIDs = ref<boolean[]>([]);

watch(event, () => {
  dateOptionIDs.value = event.value?.dateOptions?.map((v) => false) ?? [];
});

const comment = ref("");
const postDateOptionIDs = () => {
  console.log({
    comment: "",
    dateOptionIDs: dateOptionIDs.value.flatMap((v, i) =>
      v ? [event.value?.dateOptions?.[i]?.id ?? ""] : []
    ),
  });
  postApplicants({
    eventID: id,
    data: {
      comment: comment.value,
      dateOptionIDs: dateOptionIDs.value.flatMap((v, i) =>
        v ? [event.value?.dateOptions?.[i]?.id ?? ""] : []
      ),
    },
  });
};
</script>

<template>
  <div>
    <div>
      <div v-if="!event?.isConfirmed">
        <div
          v-for="dateOption in event?.dateOptions ?? []"
          :key="dateOption.id"
        >
          <v-checkbox
            v-model="dateOptionIDs"
            :value="dateOption.id"
            :label="`${format(new Date(dateOption.start), 'yyyy/MM/dd(eee)')}
              ${format(new Date(dateOption.start), 'HH:mm')} ~ ${format(
              new Date(dateOption.end),
              'HH:mm'
            )}`"
            hide-details
          >
          </v-checkbox>
        </div>

        <h2 class="text-h4">コメント</h2>
        <v-textarea
          v-model="comment"
          label="コメント"
          auto-grow
          rows="2"
        ></v-textarea>
        <v-btn @click="postDateOptionIDs" color="blue" size="x-large"
          >保存</v-btn
        >
      </div>

      <div v-if="event?.isConfirmed">
        <div>
          {{ format(new Date(event?.date?.start ?? ""), "MM/dd(eee)") }}
          {{ format(new Date(event?.date?.start ?? ""), "HH:mm") }} ~
          {{ format(new Date(event?.date?.end ?? ""), "HH:mm") }}
        </div>
      </div>
    </div>
  </div>
</template>

<style></style>
