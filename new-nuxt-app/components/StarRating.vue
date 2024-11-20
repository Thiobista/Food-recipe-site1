<template>
    <div>
      <!-- Star Rating Buttons -->
      <div class="flex items-center">
        <button
          v-for="n in 5"
          :key="n"
          class="text-3xl"
          :class="{'text-yellow-500': n <= rating}"
          @click="selectRating(n)"
        >
          <i class="fas fa-star"></i>
        </button>
      </div>
  
      <!-- Confirmation Dialog -->
      <div v-if="isDialogVisible" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
        <div class="bg-white p-5 rounded shadow-lg">
          <p class="text-lg">Are you sure you want to submit this rating?</p>
          <div class="flex justify-between mt-4">
            <button @click="confirmRating" class="bg-green-500 text-white py-2 px-4 rounded">OK</button>
            <button @click="cancelRating" class="bg-red-500 text-white py-2 px-4 rounded">Cancel</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const rating = ref(0); // Current rating value
  const isDialogVisible = ref(false); // Control dialog visibility
  
  // Handle selecting a rating
  const selectRating = (value) => {
    rating.value = value;
    isDialogVisible.value = true; // Show confirmation dialog after selecting a rating
  };
  
  // Handle confirmation (OK) action
  const confirmRating = () => {
    console.log('Rating confirmed:', rating.value);
    isDialogVisible.value = false; // Hide the dialog after confirming
  };
  
  // Handle cancellation action
  const cancelRating = () => {
    rating.value = 0; // Reset rating if user cancels
    isDialogVisible.value = false; // Hide the dialog after cancelling
  };
  </script>
  
  <style scoped>
  button {
    cursor: pointer;
  }
  </style>
  