<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { format } from "date-fns";

import {
  useGetEventsEventID,
  useGetEventsEventIDApplicants,
  useGetMe,
  usePatchEventsEventIDConfirm,
} from "../generated/api/openapi";
import { useRoute } from "vue-router";

import { mdiDotsHorizontal } from "@mdi/js";

const route = useRoute();

const id = Array.isArray(route.params.id)
  ? route.params.id[0]
  : route.params.id;

const { data: eventsAxios } = useGetEventsEventID(id);
const { data: eventsAxiosApplicants } = useGetEventsEventIDApplicants(id);
const { data: me } = useGetMe();
const { mutateAsync: postApplicants } = usePatchEventsEventIDConfirm();

const event = computed(() => eventsAxios.value?.data);

const dateID = ref("");

watch(event, () => {
  dateID.value = eventsAxios.value?.data?.dateOptions?.[0]?.id ?? "";
});

const postDateOptionIDs = () => {
  console.log(dateID.value);
  console.log(
    eventsAxiosApplicants.value?.data.filter((v) =>
      v.dateOptionIDs?.includes(dateID.value)
    )
  );
};
</script>

<template>
  <div>
    <div v-if="!event?.isConfirmed">
      <v-radio-group v-model="dateID">
        <div
          v-for="dateOption in event?.dateOptions"
          :key="dateOption.id"
          class="selectDateOption"
        >
          <v-radio :value="dateOption.id" hide-details>
            <template v-slot:label>
              {{
                `${format(new Date(dateOption.start), "yyyy/MM/dd(eee)")}
              ${format(new Date(dateOption.start), "HH:mm")} ~ ${format(
                  new Date(dateOption.end),
                  "HH:mm"
                )}`
              }}
              <v-avatar
                v-for="applicant in eventsAxiosApplicants?.data
                  .filter((v) => v.dateOptionIDs?.includes(dateOption.id))
                  .slice(0, 6)"
                :key="applicant.traqID"
                :image="`https://q.trap.jp/api/v3/public/icon/${applicant.traqID}`"
                :alt="`${applicant.traqID}'s icon'`"
                height="25"
                width="25"
              />
              <v-avatar
                v-if="
                  eventsAxiosApplicants?.data.filter((v) =>
                    v.dateOptionIDs?.includes(dateOption.id)
                  ).length &&
                  eventsAxiosApplicants?.data.filter((v) =>
                    v.dateOptionIDs?.includes(dateOption.id)
                  ).length > 6
                "
                :icon="mdiDotsHorizontal"
                height="25"
                width="25"
              />
              {{
                eventsAxiosApplicants?.data.filter((v) =>
                  v.dateOptionIDs?.includes(dateOption.id)
                ).length
              }}人
            </template>
          </v-radio>
        </div>
      </v-radio-group>
      <v-btn @click="postDateOptionIDs" color="blue" size="x-large">決定</v-btn>
    </div>
  </div>
</template>

<style>
.selectDateOption {
  display: flex;
}
</style>
