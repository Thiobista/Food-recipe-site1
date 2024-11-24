<template>
    <div class="min-h-screen bg-gray-100 p-6">
      <h1 class="text-3xl font-bold text-center text-blue-500">Contact Us</h1>
      <form class="max-w-md mx-auto mt-6 bg-white p-4 shadow-md rounded-lg">
        <div class="mb-4">
          <label for="name" class="block text-gray-700">Name:</label>
          <input type="text" id="name" class="w-full border rounded p-2" />
        </div>
        <div class="mb-4">
          <label for="email" class="block text-gray-700">Email:</label>
          <input type="email" id="email" class="w-full border rounded p-2" />
        </div>
        <div class="mb-4">
          <label for="message" class="block text-gray-700">Message:</label>
          <textarea id="message" class="w-full border rounded p-2"></textarea>
        </div>
        <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
          Submit
        </button>
      </form>
    </div>
  </template>
  