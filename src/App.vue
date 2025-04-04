<script setup>
import { ref } from "vue";
import { Sheet } from "bottom-sheet-vue3";

const visible = ref(false);
</script>

<template>
  <button @click="visible = true">Click to show sheet</button>
  <div class="flex justify-center">
    <div class="max-w-x min-w-[200px]">
      <input
        class="rounded-full shadow-lg p-2 w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-lg border border-slate-200 px-4 py-4 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300  focus:shadow"
        placeholder="Where to?"
        @click="visible = true"
      />
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
        '--bottom-sheet-backdrop-background-color' : 'rgba(0, 0, 0, 0)',
      }"
    >
      Hello this is Sheet Content
    </Sheet>
  </div>
</template>
