<template>
    <form @submit.prevent="submitRecipe" class="space-y-6">
      <div class="flex flex-col">
        <label for="title" class="font-medium">Recipe Title</label>
        <input type="text" id="title" v-model="title" required class="px-4 py-2 rounded-lg border border-gray-300">
      </div>
  
      <div class="flex flex-col">
        <label for="description" class="font-medium">Description</label>
        <textarea id="description" v-model="description" required class="px-4 py-2 rounded-lg border border-gray-300"></textarea>
      </div>
  
      <div class="flex justify-between">
        <button type="submit" class="bg-yellow-500 text-white px-6 py-3 rounded-full text-lg font-semibold hover:bg-yellow-400 transition">Submit Recipe</button>
      </div>
    </form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        title: '',
        description: ''
      }
    },
    methods: {
      submitRecipe() {
        // Handle form submission
      }
    }
  }
  </script>
  