<script setup>
import { ref, watch } from "vue";
import { Sheet } from "bottom-sheet-vue3";
import InputFields from "./components/InputFields.vue"
import LocationSelector from './components/LocationSelector.vue'
import { useBottomSheet } from '@/stores/useBottomSheet'

const visible = ref(false);


const location_selector = useBottomSheet();


watch(visible, () =>{
  if(visible.value && location_selector.location_visible){
    location_selector.location_visible = false;
  }
})

function blele(){
  console.log(location_selector.location_visible);
}

</script>

<template>
  <button @click="blele">1123123</button>
  <div
    v-if="!visible"
    @click="visible = 'true'"
    class="absolute left-0 right-0 bottom-0 bg-white rounded-t-3xl shadow-xl transition-transform duration-300 ease-in-out transfor"
  >
    <div class="h-[200px] flex flex-col cursor-pointer overflow-hidden fixed bottom-0 bg-white w-full z-100">
      <div class="flex justify-center mt-3">
              <div class="w-12 h-1 bg-gray-300 rounded-full"></div>
      </div>
       <InputFields />
    </div>
  </div>
  <div>
    <Sheet
      v-model:visible="visible"
      :style="{
        '--bottom-sheet-max-width': '512px',
        '--bottom-sheet-min-width': '100%',
        '--bottom-sheet-max-height' : '100%',
        '--bottom-sheet-header-bar-background-color': 'rgb(209,213,219)',
        '--bottom-sheet-backdrop-background-color': 'rgba(0, 0, 0, 0)',
      }"
    >

      <InputFields v-if="!location_selector.location_visible"/> 
      <LocationSelector v-if="location_selector.location_visible"/>
      <button @click="blele">1123123</button>
    <template #header>
      <div class="mt-3 w-12 h-1 bg-gray-300 rounded-full" :visible='visible'></div>
    </template>
    </Sheet>
  </div>
</template>
