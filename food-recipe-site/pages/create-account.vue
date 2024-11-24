<template>
  <div class="min-h-screen bg-gradient-to-r from-teal-400 to-blue-500 flex items-center justify-center">
    <!-- Card Container -->
    <div class="w-full max-w-md bg-white p-8 rounded-lg shadow-xl transform transition-all duration-500 hover:scale-105">
      <!-- Title -->
      <h1 class="text-4xl font-semibold text-center text-gray-800 mb-6">Create Your Account</h1>
      
      <!-- Form -->
      <form @submit.prevent="submitForm">
        <div class="space-y-6">
          <!-- Full Name Input -->
          <div>
            <div class="relative">
              <input
                type="text"
                v-model="fullName"
                placeholder="Full Name"
                :class="{'border-red-500': errors.fullName}"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
              <span v-if="errors.fullName" class="text-red-500 text-sm">{{ errors.fullName }}</span>
            </div>
          </div>

          <!-- Email Input -->
          <div>
            <div class="relative">
              <input
                type="email"
                v-model="email"
                placeholder="Email Address"
                :class="{'border-red-500': errors.email}"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
              <span v-if="errors.email" class="text-red-500 text-sm">{{ errors.email }}</span>
            </div>
          </div>

          <!-- Password Input -->
          <div>
            <div class="relative">
              <input
                type="password"
                v-model="password"
                placeholder="Password"
                :class="{'border-red-500': errors.password}"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
              <span v-if="errors.password" class="text-red-500 text-sm">{{ errors.password }}</span>
            </div>
          </div>

          <!-- Submit Button -->
          <div>
            <button
              type="submit"
              class="w-full bg-teal-500 text-white py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-teal-600"
            >
              Create Account
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
// Reactive variables for form data
import { ref } from 'vue';

const fullName = ref('');
const email = ref('');
const password = ref('');
const errors = ref({});

// Function to handle form submission
const submitForm = () => {
  // Reset errors
  errors.value = {};

  // Validate form
  if (!fullName.value) {
    errors.value.fullName = 'Full name is required';
  }
  if (!email.value) {
    errors.value.email = 'Email is required';
  } else if (!/\S+@\S+\.\S+/.test(email.value)) {
    errors.value.email = 'Email is invalid';
  }
  if (!password.value) {
    errors.value.password = 'Password is required';
  }

  // If no errors, emit initials and navigate
  if (Object.keys(errors.value).length === 0) {
    const initials = fullName.value.split(' ').map(word => word[0].toUpperCase()).join('');
    
    // You can emit to a parent component or handle routing
    // Example: use a store or navigate to another page

    // Use Nuxt's useRouter() to redirect to the homepage
    const router = useRouter();
    router.push('/');
  }
};
</script>

<style scoped>
/* You can customize the styles if needed */
</style>
