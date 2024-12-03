<template>
    <div class="relative min-h-screen bg-cover bg-center flex items-center justify-center">
      <Header />
      <!-- Blurred Backghround Layer -->
      <div 
        class="absolute inset-0 bg-cover bg-center blur-sm" 
        style="background-image: url('images/chief.jpg');"
      ></div>
  
      <!-- Content Layer -->
      <div class="relative w-full max-w-md bg-white p-8 rounded-lg shadow-xl transform transition-all duration-500 hover:scale-105">
        <!-- Title -->
        <h1 class="text-4xl font-semibold text-center text-gray-800 mb-6">Login </h1>
        
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
                :disabled="loading"
                type="submit"
                class="w-full bg-teal-500 text-white py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-teal-600 disabled:opacity-50"
              >
                {{ loading ? 'login to your Account...' : 'Login' }}
              </button>
            </div>
          </div>
        </form>
  
        <!-- Continue Without Account Button -->
        <div class="mt-4 text-center">
          <button
            @click="continueWithoutAccount"
            class="text-teal-500 underline hover:text-teal-700 transition duration-200"
          >
            Continue Without Account
          </button>
        </div>
  
        <!-- Login Button -->
        <!-- <div class="mt-2 text-center">
          <button
            @click="goToLogin"
            class="text-teal-500 underline hover:text-teal-700 transition duration-200"
          >
            Already have an account? Login
          </button>
        </div> -->
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import axios from 'axios';
  
  const router = useRouter();
  const fullName = ref('');
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
        // Navigate to the authorized page
        alert("Login successful!");
        router.push("/authorized");
      }
    } catch (error) {
      // Show error alert
      if (error.response && error.response.status === 401) {
        alert("Invalid email or password.");
      } else {
        alert("An unexpected error occurred. Please try again later.");
      }
    } finally {
      loading.value = false;
    }
  }
};

  
  const continueWithoutAccount = () => {
    router.push('/');
  };
  
  const goToLogin = () => {
    router.push('/login'); // Navigate to the login page
  };
  const login = async () => {
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

      // Handle success
      alert(response.data.message);
      // Store the user ID if needed (e.g., in localStorage)
    } catch (error) {
      // Handle errors
      alert(error.response.data || "Login failed");
    } finally {
      loading.value = false;
    }
  }
};

  </script>
  