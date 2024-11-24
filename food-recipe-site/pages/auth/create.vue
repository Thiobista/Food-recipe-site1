<template>
  <div class="min-h-screen flex flex-col justify-center items-center bg-gray-100">
    <div class="bg-white shadow-lg rounded-lg p-8 max-w-sm w-full">
      <h1 class="text-2xl font-bold mb-4 text-center">Create Account</h1>

      <!-- Input for Name -->
      <input
        v-model="name"
        type="text"
        placeholder="Name"
        class="w-full p-2 mb-4 border rounded-lg"
      />

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

      <!-- Sign Up Button -->
      <button
        @click="handleSignUp"
        class="w-full bg-yellow-500 text-white py-2 rounded-lg font-bold"
      >
        Sign Up
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// Reactive variables for form inputs
const name = ref('');
const email = ref('');
const password = ref('');

// Function to handle form submission
const handleSignUp = () => {
  if (!name.value || !email.value || !password.value) {
    alert('Please fill in all the fields.');
    return;
  }

  // Example sign-up logic
  console.log('Name:', name.value);
  console.log('Email:', email.value);
  console.log('Password:', password.value);

  alert(`Account successfully created for ${name.value}!`);
};
</script>
