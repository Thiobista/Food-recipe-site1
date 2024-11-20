<template>
    <div class="p-6 bg-white shadow-lg rounded-lg">
      <h2 class="text-2xl font-bold mb-4">Enter Payment Details</h2>
      <form @submit.prevent="submitPayment">
        <div class="mb-4">
          <label for="name" class="block text-gray-700">Name</label>
          <input type="text" id="name" v-model="name" class="w-full p-2 border rounded-lg" required />
        </div>
        <div class="mb-4">
          <label for="email" class="block text-gray-700">Email</label>
          <input type="email" id="email" v-model="email" class="w-full p-2 border rounded-lg" required />
        </div>
        <div class="mb-4">
          <label for="amount" class="block text-gray-700">Amount</label>
          <input type="number" id="amount" v-model="amount" class="w-full p-2 border rounded-lg" required />
        </div>
        <button type="submit" class="bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600">
          Submit
        </button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        name: "",
        email: "",
        amount: 0,
      };
    },
    methods: {
      submitPayment() {
        this.$toast.success("Payment submitted successfully!");
        // Add your Chapa API integration logic here
      },
    },
  };
  </script>
  
  <style scoped>
  input {
    transition: border 0.3s ease;
  }
  input:focus {
    border-color: #38a169;
  }
  </style>
  