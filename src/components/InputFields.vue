<script setup>
import { MapPin, Navigation } from "lucide-vue-next";
</script>

<template>
	<div class="p-4">
		<h2 class="text-xl font-bold">Book a Ride</h2>
	</div>
	<div class="flex items-center justify-center">
		<div class="flex flex-col items-center relative pl-4">
			<MapPin
				size="12"
				color="#ff8000"
				class="absolute translate-y-[-14px] translate-x-[-1.5px]"
			/>

			<div
				class="w-1 h-13 border-l-2 border-dashed border-gray-300"
			></div>

			<Navigation
				size="12"
				color="#ff8000"
				class="absolute translate-y-[52px] translate-x-[-2.7px]"
			/>
		</div>
		<div class="flex flex-col p-4 w-full gap-4">
			<div class="w-full min-w-[200px]">
				<input
					class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-4 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
					placeholder="Pickup location"
				/>
			</div>
			<div class="w-full min-w-[200px]">
				<input
					class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-4 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
					placeholder="Destination"
				/>
			</div>
		</div>
	</div>
</template>
