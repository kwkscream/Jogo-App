<script setup>
import { ref } from "vue";
import { Sheet } from "bottom-sheet-vue3";
import InputFields from "./components/InputFields.vue"


const visible = ref(false);
</script>

<template>
  <div class="flex justify-center">
    <div class="max-w-x min-w-[200px] p-4">
      <input
        class="rounded-full shadow-lg w-full bg-transparent placeholder:text-slate-700 text-neutral-900 text-balance border border-slate-300 px-3 py-3 transition duration-300 ease focus:outline-none focus:border-slate-600 hover:border-slate-300 focus:shadow focus-visible:ring-0 focus-visible:ring-offset-0"
        placeholder="Where to?"
        @click="visible = true"
      />
    </div>
  </div>

  <div
    class="absolute left-0 right-0 bottom-0 bg-white rounded-t-3xl shadow-xl transition-transform duration-300 ease-in-out transfor"
    @click="visible = 'true'"
  >
    <div class="h-6 w-full flex items-center justify-center cursor-pointer">
      <div class="w-12 h-1 bg-gray-300 rounded-full"></div>
    </div>
  </div>
  <div>
    <Sheet
      v-model:visible="visible"
      :style="{
        '--bottom-sheet-max-width': '512px',
        '--bottom-sheet-min-width': '70%',
        '-bottom-sheet-max-height': '80%',
        '--bottom-sheet-header-bar-background-color': 'rgb(209,213,219)',
        '--bottom-sheet-backdrop-background-color': 'rgba(0, 0, 0, 0)',
      }"
    >
      <InputFields></InputFields>
    <template #header>
      <div class="mt-3 w-12 h-1 bg-gray-300 rounded-full"></div>
    </template>
    </Sheet>
  </div>
</template>
