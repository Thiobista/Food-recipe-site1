<template>
  <div class="relative min-h-screen bg-cover bg-center flex items-center justify-center">
    <!-- Blurred Background Layer -->
    <div 
      class="absolute inset-0 bg-cover bg-center blur-sm" 
      style="background-image: url('images/chief.jpg');"
    ></div>

    <!-- Content Layer -->
    <div class="relative w-full max-w-md bg-white p-8 rounded-lg shadow-xl transform transition-all duration-500 hover:scale-105">
      <!-- Title -->
      <h1 class="text-4xl font-semibold text-center text-gray-800 mb-6">Login</h1>
      
      <!-- Form -->
      <form @submit.prevent="submitForm">
        <div class="space-y-6">
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
              :disabled="loading"
              type="submit"
              class="w-full bg-teal-500 text-white py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-teal-600 disabled:opacity-50"
            >
              {{ loading ? 'Logging in...' : 'Login' }}
            </button>
          </div>
        </div>
      </form>

      <!-- Continue Without Account Button -->
      <div class="mt-4 text-center">
        <button
          @click="goToLogin"
          class="w-full bg-gray-100 text-gray-700 py-2 px-4 rounded-lg shadow-md hover:bg-gray-200 transform hover:scale-105 transition duration-300"
        >
          Already have an account? Login
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { useToast } from 'vue-toastification';

const router = useRouter();
const toast = useToast();
const email = ref('');
const password = ref('');
const errors = ref({});
const loading = ref(false);

const submitForm = async () => {
  errors.value = {};
  if (!email.value) errors.value.email = "Email is required";
  if (!password.value) errors.value.password = "Password is required";

  if (Object.keys(errors.value).length === 0) {
    try {
      loading.value = true;

      const response = await axios.post("http://localhost:8081/login", {
        email: email.value,
        password: password.value,
      });

      if (response.status === 200) {
        const token = response.data.token; // Assuming the token is in response.data.token
        localStorage.setItem("token", token); // Store the token in localStorage
        toast.success(response.data.message || "Login successful!");
        router.push("/authorized");
      }
    } catch (error) {
      if (error.response && error.response.status === 401) {
        toast.error(error.response.data || "Invalid credentials. Please try again.");
      } else {
        toast.error("An unexpected error occurred. Please try again later.");
      }
    } finally {
      loading.value = false;
    }
  }
};

const continueWithoutAccount = () => {
  router.push('/');
};
</script>
