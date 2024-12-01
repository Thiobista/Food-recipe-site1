<template>
  <div class="container mx-auto">
    <!-- Header Section -->
    <header class="fixed top-0 left-0 w-full bg-white shadow-md z-50">
      <Header />
    </header>

    <!-- Main Content -->
    <div class="pt-20"> <!-- Adds padding to avoid overlap with the fixed header -->
      <h1 class="text-4xl font-semibold text-center text-gray-800 mb-8">Profile Management</h1>
      <div class="max-w-md mx-auto bg-white p-8 rounded-lg shadow-lg">
        <!-- Profile Image Section -->
        <div class="flex flex-col items-center mb-6">
          <label class="text-lg font-semibold text-gray-700">Profile Image</label>
          <img :src="profileImage" alt="Profile" class="w-24 h-24 rounded-full border-4 border-yellow-500 mt-4 mb-4" />
          <input
            type="file"
            @change="updateProfileImage"
            class="px-4 py-2 text-sm text-gray-800 rounded-md bg-gray-200 cursor-pointer hover:bg-gray-300 transition-colors"
          />
        </div>

        <!-- Username Section -->
        <div class="mb-6">
          <label class="block text-lg font-medium text-gray-700">Username</label>
          <input
            v-model="username"
            type="text"
            class="w-full px-4 py-3 mt-2 text-gray-800 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:border-transparent transition-all"
            placeholder="Enter new username"
          />
        </div>

        <!-- Password Section -->
        <div class="mb-6">
          <label class="block text-lg font-medium text-gray-700">Password</label>
          <input
            v-model="password"
            type="password"
            class="w-full px-4 py-3 mt-2 text-gray-800 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:border-transparent transition-all"
            placeholder="Enter new password"
          />
        </div>

        <!-- Save Button -->
        <button
          @click="saveChanges"
          class="w-full bg-yellow-500 text-white px-6 py-3 rounded-md shadow-lg hover:bg-yellow-600 focus:outline-none transition-colors"
        >
          Save Changes
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
// Import the useRouter from Nuxt 3
import { ref } from 'vue';
import { useRouter } from 'vue-router';

// Define reactive variables for profile image, username, and password
const profileImage = ref('/images/profile.jpg');
const username = ref('');
const password = ref('');

// Router instance to navigate to other pages
const router = useRouter();

// Method to update the profile image
const updateProfileImage = (event) => {
  const file = event.target.files[0];
  if (file) {
    profileImage.value = URL.createObjectURL(file);
  }
};

// Method to save changes and navigate to the home page
const saveChanges = () => {
  console.log('Profile updated:', username.value, password.value);

  // Navigate to the home page (index.vue) using the Nuxt router
  router.push('/'); // Navigate to the home page
};
</script>

<style scoped>
/* Global Styling */
body {
  background-color: #f9fafb;
  font-family: 'Inter', sans-serif;
}

/* Styling for inputs and button */
input[type="text"],
input[type="password"],
input[type="file"] {
  transition: all 0.3s ease;
}

/* Focus styles for inputs */
input:focus {
  box-shadow: 0 0 0 2px rgba(234, 179, 8, 0.6);
}

/* Button Hover */
button:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style>
