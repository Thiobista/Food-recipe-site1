<template>
    <header>
      <nav class="bg-gradient-to-r from-yellow-400 via-yellow-500 to-yellow-600 p-4 flex justify-between items-center shadow-xl rounded-b-lg fixed top-0 left-0 w-full z-20">
        <!-- Back Arrow -->
        <button 
          @click="goBack" 
          class="text-white font-semibold text-xl hover:text-gray-300 focus:outline-none flex items-center space-x-2"
        >
          <span>←</span>
          <span>Back</span>
        </button>
  
        <!-- Menu Button -->
        <button
          @click="toggleMenu"
          class="text-white font-semibold text-xl hover:text-gray-300 focus:outline-none"
        >
          ☰ Menu
        </button>
  
        <!-- Menu Contents -->
        <transition name="slide-fade">
          <div
            v-if="menuOpen"
            class="fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-30 overflow-y-auto"
          >
            <div class="p-4 bg-yellow-500 text-white font-bold">
              Menu
              <button @click="toggleMenu" class="absolute top-4 right-4 text-xl font-bold">
                ✕
              </button>
            </div>
            <ul class="space-y-4 mt-4 px-4">
              <li
                v-for="(item, index) in menuItems"
                :key="index"
                @click="goToPage(item.route)"
                class="p-3 hover:bg-yellow-100 flex items-center space-x-3 cursor-pointer"
              >
                <span>{{ item.name }}</span>
              </li>
            </ul>
          </div>
        </transition>
      </nav>
    </header>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  
  // State Variables
  const menuOpen = ref(false);
  
  // Menu Items
  const menuItems = ref([
    { name: 'Profile', route: '/profile' ,img: '/images/profile.jpg', alt: 'Profile'},
    { name: 'Create Recipe', route: '/create-recipe' },
    { name: 'My Recipes', route: '/my-recipes' },
    { name: 'Bookmarks', route: '/bookmark' },
    { name: 'Logout', route: '/logout' },
  ]);
  
  // Router Instance
  const router = useRouter();
  
  // Methods
  const toggleMenu = () => (menuOpen.value = !menuOpen.value);
  const goToPage = (route) => router.push(route);
  const goBack = () => router.back(); // Navigate to the previous page
  </script>
  
  <style scoped>
  /* Add specific styles for the header here */
  </style>
  