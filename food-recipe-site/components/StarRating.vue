<template>
  <div class="p-6">
    <!-- Star Rating Buttons -->
    <div class="flex items-center justify-center space-x-2">
      <button
        v-for="n in 5"
        :key="n"
        class="text-3xl transition duration-300"
        :class="{'text-yellow-500': n <= rating, 'text-gray-400': n > rating}"
        @click="selectRating(n)"
        aria-label="Rate {{ n }} star"
      >
        <i class="fas fa-star"></i>
      </button>
    </div>

    <!-- Confirmation Dialog -->
    <div
      v-if="isDialogVisible"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
    >
      <div class="bg-white p-5 rounded shadow-lg max-w-md mx-auto text-center">
        <p class="text-lg font-medium">Are you sure you want to submit this rating?</p>
        <div class="flex justify-center space-x-4 mt-4">
          <button
            @click="confirmRating"
            class="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-600 transition"
          >
            OK
          </button>
          <button
            @click="cancelRating"
            class="bg-red-500 text-white py-2 px-4 rounded hover:bg-red-600 transition"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// Reactive variables
const rating = ref(0); // Current rating value
const isDialogVisible = ref(false); // Control dialog visibility

// Handle selecting a rating
const selectRating = (value) => {
  rating.value = value;
  isDialogVisible.value = true; // Show confirmation dialog after selecting a rating
};

// Handle confirmation (OK) action
const confirmRating = () => {
  console.log("Rating confirmed:", rating.value);
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

button:hover {
  transform: scale(1.1);
}
</style>
