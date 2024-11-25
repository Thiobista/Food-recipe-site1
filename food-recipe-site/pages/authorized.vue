<template>
    <div>
      <!-- Navbar Section -->
      <header>
        <nav class="bg-gradient-to-r from-yellow-400 via-yellow-500 to-yellow-600 p-4 flex justify-between items-center shadow-xl rounded-b-lg fixed top-0 left-0 w-full z-20">
          <!-- Menu Button (Top Left) -->
          <div>
            <button
              @click="toggleMenu"
              class="text-white font-semibold text-xl hover:text-gray-300 focus:outline-none"
            >
              ☰ Menu
            </button>
          </div>
  
          <!-- Menu Contents -->
          <transition name="slide-fade">
            <div
              v-if="menuOpen"
              class="fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-30 overflow-y-auto transition-transform transform"
            >
              <div class="p-4 bg-yellow-500 text-white font-bold">
                Menu
                <button
                  @click="toggleMenu"
                  class="absolute top-4 right-4 text-xl font-bold"
                >
                  ✕
                </button>
              </div>
              <ul class="space-y-4 mt-4 px-4">
                <li
                  @click="goToPage('/profile')"
                  class="p-3 hover:bg-yellow-100 flex items-center space-x-3 cursor-pointer"
                >
                  <img src="/images/profile.jpg" alt="Profile" class="w-8 h-8 rounded-full border-2 border-yellow-500" />
                  <span>Profile</span>
                </li>
                <li
                  @click="goToPage('/create-recipe')"
                  class="p-3 hover:bg-yellow-100 cursor-pointer"
                >
                  Create Recipe
                </li>
                <li
                  @click="goToPage('/my-recipes')"
                  class="p-3 hover:bg-yellow-100 cursor-pointer"
                >
                  My Recipes
                </li>
                <li
                  @click="goToPage('/bookmarks')"
                  class="p-3 hover:bg-yellow-100 cursor-pointer"
                >
                  Bookmarks
                </li>
                <li
                  @click="goToPage('/logout')"
                  class="p-3 hover:bg-yellow-100 text-red-500 cursor-pointer"
                >
                  Logout
                </li>
              </ul>
            </div>
          </transition>
  
          <!-- Create Account Button (Top Right) -->
          <!-- <div class="ml-auto">
            <button
              class="bg-yellow-500 text-white py-2 px-6 rounded-lg shadow-lg hover:bg-yellow-600 transition duration-300"
              @click="createAccount"
            >
              Create Account
            </button>
          </div> -->
        </nav>
      </header>
   <!-- Hero Section -->
   <section
        class="relative bg-cover bg-center h-[60vh] rounded-lg overflow-hidden mt-16"
        style="background-image: url('/images/hero.jpg');"
      >
        <div class="absolute inset-0 bg-black opacity-60"></div>
        <div class="relative z-10 text-center text-white py-24 px-6">
          <h1 class="text-4xl font-extrabold mb-4 leading-tight">Welcome to Delicious Recipes</h1>
          <p class="text-xl mb-6 font-light">Explore, Create, and Share your favorite recipes.</p>
          <button
            class="bg-yellow-500 text-white py-3 px-8 rounded-full text-lg shadow-lg hover:bg-yellow-600 transform hover:scale-105 transition duration-300 ease-in-out"
          @click="scrollToRecipes"
            >
            Start Cooking 
          </button>
        </div>
      </section>
      <!-- Main Content -->
      <main class="container mx-auto py-12">
        <SearchFilterBar />
  
        <!-- Recipe Cards Grid -->
        <div class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-8">
          <RecipeCard
            v-for="n in 6"
            :key="n"
            image="/images/salad.jpg"
            title="Delicious Recipe"
            description="A quick preview of the recipe details."
            @rate="rateRecipe"
            @comment="commentOnRecipe"
            @bookmark="bookmarkRecipe"
            @buy="goToPaymentPage"
          />
        </div>
      </main>
      <!-- Footer -->
      <Footer />
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  
  import SearchFilterBar from '~/components/SearchFilterBar.vue';
  import RecipeCard from '~/components/RecipeCard.vue';
  import Footer from '~/components/Footer.vue';
  
  const menuOpen = ref(false);
  const bookmarkedRecipes = ref([]);
  const router = useRouter();
  
  const toggleMenu = () => {
    menuOpen.value = !menuOpen.value;
  };
  
  const goToPage = (route) => {
    toggleMenu(); // Close the menu when navigating
    router.push(route);
  };
  
  const createAccount = () => {
    router.push('/create-account');
  };
  //  const rateRecipe = (rating) => {
  //    console.log(Rated this recipe ${rating} stars);
  // };
  
  const commentOnRecipe = () => {
    console.log('Write your comment!');
  };
  
  const bookmarkRecipe = (recipe) => {
    const exists = bookmarkedRecipes.value.some((r) => r.id === recipe.id);
    if (!exists) {
      bookmarkedRecipes.value.push(recipe);
      console.log('Recipe added to bookmarks!');
      router.push('/bookmarks');  // Navigate to bookmarks page
    } else {
      console.log('Recipe is already bookmarked.');
    }
  };
  
  
  const goToPaymentPage = () => {
    router.push('/payment');
  };
  const scrollToRecipes = () => {
    const recipesSection = document.querySelector('main');
    if (recipesSection) {
      recipesSection.scrollIntoView({ behavior: 'smooth' });
    }
  };
  </script>
  
  <style scoped>
  /* Transition for the menu sliding effect */
  .slide-fade-enter-active,
  .slide-fade-leave-active {
    transition: transform 0.3s ease, opacity 0.3s ease;
  }
  
  .slide-fade-enter-from {
    transform: translateX(-100%);
    opacity: 0;
  }
  
  .slide-fade-leave-to {
    transform: translateX(-100%);
    opacity: 0;
  }
  
  /* General styling for the menu */
  .fixed {
    background-color: #ffffff;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    z-index: 30;
  }
  
  button:hover {
    transform: scale(1.05);
    transition: transform 0.3s ease-in-out;
  }
  
  /* Menu button and dropdown menu responsiveness */
  @media (max-width: 768px) {
    .menu-button {
      font-size: 1.25rem;
    }
  
    .dropdown-menu[aria-expanded="true"] {
      width: 100%;
      opacity: 1;
    }
  }
  
  /* Hover and transition effects for buttons */
  button:hover {
    transform: scale(1.05);
    transition: transform 0.3s ease-in-out;
  }
  
  /* Dropdown menu item hover effects */
  li:hover {
    background-color: #f9e2a0; /* Lighter yellow for hover */
    transform: scale(1.05); /* Slightly larger effect */
  }
  
  /* Hover effects for profile image */
  li img {
    border: 2px solid #fbbf24; /* Yellow border on hover */
  }
  
  /* Hero Section styling */
  .hero-section {
    background: url('/images/salad.jpg') center/cover no-repeat;
    height: 60vh;
    position: relative;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    background-color: #f9f9f9; /* Neutral fallback color */
    background-image: url('/images/salad.jpg');
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    height: 60vh;
    position: relative;
  }
  /* Lazy load styling */
  img[loading="lazy"] {
    filter: blur(10px);
    transition: filter 0.3s ease-in-out;
  }
  
  img[loading="lazy"]:loaded {
    filter: blur(0);
  }
  .hero-section h1 {
    font-size: 3rem;
    font-weight: bold;
    margin-bottom: 1rem;
  }
  
  .hero-section p {
    font-size: 1.25rem;
    margin-bottom: 1.5rem;
  }
  
  /* Recipe Card styling */
  .recipe-card {
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
  }
  
  .recipe-card:hover {
    transform: scale(1.05);
    box-shadow: 0 6px 18px rgba(0, 0, 0, 0.2);
  }
  
  .recipe-card img {
    width: 100%;
    height: 200px;
    object-fit: cover;
    border-radius: 8px;
    width: 100%;
    height: 200px; /* Adjust height for uniformity */
    object-fit: cover; /* Ensures the image is centered and fills the space */
    border-radius: 8px;
  }
  /* For smaller screens */
  @media (max-width: 768px) {
    .recipe-card img {
      height: 150px; /* Reduce height for mobile */
    }
  }
  /* Footer styling */
  footer {
    background-color: #f8f8f8;
    padding: 2rem;
    text-align: center;
    color: #333;
  }
  
  footer a {
    text-decoration: none;
    color: #333;
  }
  
  footer a:hover {
    color: #ff8c00;
  }
  
  /* Add more spacing and modernize typography */
  body {
    font-family: 'Roboto', sans-serif;
  }
  
  /* Custom transition for the menu */
  .menu-item-enter-active, .menu-item-leave-active {
    transition: opacity 0.5s ease;
  }
  
  .menu-item-enter, .menu-item-leave-to {
    opacity: 0;
  }
  
  /* Improved shadow effect for the menu */
  .dropdown-menu {
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  }
  
  /* Add a smooth animation to the profile image */
  li img {
    transition: transform 0.3s ease;
  }
  
  li:hover img {
    transform: rotate(360deg); /* Rotate the profile image on hover */
  }
  
  </style>
  