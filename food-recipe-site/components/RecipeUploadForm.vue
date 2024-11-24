<template>
  <div class="p-8 bg-gray-100 rounded-lg shadow-md max-w-md mx-auto">
    <form @submit.prevent="submitRecipe" class="space-y-6">
      <!-- Recipe Title -->
      <div class="flex flex-col">
        <label for="title" class="font-medium">Recipe Title</label>
        <input
          type="text"
          id="title"
          v-model="title"
          required
          placeholder="Enter your recipe title"
          class="px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500"
        />
      </div>

      <!-- Description -->
      <div class="flex flex-col">
        <label for="description" class="font-medium">Description</label>
        <textarea
          id="description"
          v-model="description"
          required
          placeholder="Write a brief description of your recipe"
          class="px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500"
        ></textarea>
      </div>

      <!-- Submit Button -->
      <div class="flex justify-between">
        <button
          type="submit"
          class="bg-yellow-500 text-white px-6 py-3 rounded-full text-lg font-semibold hover:bg-yellow-400 transition"
        >
          Submit Recipe
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
// Import required composables
import { ref } from "vue";
import { useRouter } from "nuxt/app"; // Nuxt 3-specific router composable

// Reactive variables for form data
const title = ref("");
const description = ref("");

// Router instance for navigation after form submission
const router = useRouter();

// Form submission logic
const submitRecipe = async () => {
  // Mock form submission logic
  console.log("Recipe Submitted:", {
    title: title.value,
    description: description.value,
  });

  // Perform any API calls using $fetch or navigate to another route
  await new Promise((resolve) => setTimeout(resolve, 500)); // Simulate API delay

  // Redirect after submission (optional)
  router.push("/success"); // Adjust this to your desired route
};
</script>

<style scoped>
/* Form container styles */
form {
  background-color: #ffffff;
  padding: 1.5rem;
  border-radius: 0.75rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

/* Focus effects */
input:focus,
textarea:focus {
  border-color: #facc15; /* Yellow border */
  box-shadow: 0 0 0 3px rgba(250, 204, 21, 0.5); /* Focus outline */
}

/* Button hover effect */
button:hover {
  background-color: #eab308; /* Slightly darker yellow */
}
</style>
