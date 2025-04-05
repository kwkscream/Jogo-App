import { defineStore } from 'pinia'

export const useBottomSheet = defineStore('counter', {
  state: () => ({
    location_visible: false
  }),
  actions: {
    changeLocationVisible(){
      this.location_visible = !this.location_visible;
    }
  }
})
