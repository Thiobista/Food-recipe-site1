<template>
  <div class="min-h-screen flex flex-col justify-center items-center bg-gray-100">
    <div class="bg-white shadow-lg rounded-lg p-8 max-w-sm w-full">
      <h1 class="text-2xl font-bold mb-4 text-center">Login</h1>

      <!-- Input for Email -->
      <input
        v-model="email"
        type="email"
        placeholder="Email"
        class="w-full p-2 mb-4 border rounded-lg"
      />

      <!-- Input for Password -->
      <input
        v-model="password"
        type="password"
        placeholder="Password"
        class="w-full p-2 mb-4 border rounded-lg"
      />

      <!-- Login Button -->
      <button
        @click="handleLogin"
        class="w-full bg-yellow-500 text-white py-2 rounded-lg font-bold"
      >
        Login
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// Reactive variables for email and password
const email = ref('');
const password = ref('');

// Function to handle login
const handleLogin = () => {
  if (!email.value || !password.value) {
    alert('Please fill in both fields.');
    return;
  }

  // Example login logic
  console.log('Email:', email.value);
  console.log('Password:', password.value);

  alert('Login successful!');
};
</script>
