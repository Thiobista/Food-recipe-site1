<template>
  <div>
    <!-- Navbar Section -->
    <header>
      <nav class="bg-gradient-to-r from-yellow-400 via-yellow-500 to-yellow-600 p-4 flex justify-between items-center shadow-xl rounded-b-lg fixed top-0 left-0 w-full z-20">
        <!-- Menu Button (Top Left) -->
        <div class="relative">
          <button
            @click="toggleMenu"
            class="text-white font-semibold text-xl hover:text-gray-300 focus:outline-none"
          >
            ☰ Menu
          </button>

          <!-- Dropdown Menu -->
          <div
            v-if="menuOpen"
            class="absolute bg-white shadow-xl rounded-lg mt-2 w-48 z-30 transition-all duration-300 ease-in-out"
          >
            <ul class="space-y-2">
              <li
                @click="goToPage('/profile')"
                class="p-3 hover:bg-yellow-100 flex items-center space-x-3 transition-all hover:scale-105 cursor-pointer"
              >
                <img src="/path/to/profile-image.jpg" alt="Profile" class="w-8 h-8 rounded-full border-2 border-yellow-500" />
                <span>Profile</span>
              </li>
              <li
                @click="goToPage('/create-recipe')"
                class="p-3 hover:bg-yellow-100 transition-all hover:scale-105 cursor-pointer"
              >
                Create Recipe
              </li>
              <li
                @click="goToPage('/my-recipes')"
                class="p-3 hover:bg-yellow-100 transition-all hover:scale-105 cursor-pointer"
              >
                My Recipes
              </li>
              <li
                @click="goToPage('/bookmarks')"
                class="p-3 hover:bg-yellow-100 transition-all hover:scale-105 cursor-pointer"
              >
                Bookmarks
              </li>
              <li
                @click="goToPage('/logout')"
                class="p-3 hover:bg-yellow-100 text-red-500 transition-all hover:scale-105 cursor-pointer"
              >
                Logout
              </li>
            </ul>
          </div>
        </div>

        <!-- Create Account Button (Top Right) -->
        <div class="auth-container ml-auto">
          <button
            class="bg-blue-500 text-white py-2 px-6 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-blue-600"
            @click="createAccount"
          >
            Create Account
          </button>
        </div>
      </nav>
    </header>

    <!-- Hero Section -->
    <section class="relative bg-cover bg-center h-[60vh] rounded-lg overflow-hidden mt-16" style="background-image: url('/images/salad.jpg');">
      <div class="absolute inset-0 bg-black opacity-60"></div>
      <div class="relative z-10 text-center text-white py-24 px-6">
        <h1 class="text-4xl font-extrabold mb-4 leading-tight">Welcome to Delicious Recipes</h1>
        <p class="text-xl mb-6 font-light">Explore, Create, and Share your favorite recipes.</p>
        <button class="bg-yellow-500 text-white py-3 px-8 rounded-full text-lg shadow-lg hover:bg-yellow-600 transform hover:scale-105 transition duration-300 ease-in-out">
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
          image="/images/enjera.jpg"
          title="Delicious enjera"
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

<script>
import SearchFilterBar from '~/components/SearchFilterBar.vue';
import RecipeCard from '~/components/RecipeCard.vue';
import Footer from '~/components/Footer.vue';
import StarRating from '~/components/StarRating.vue';

export default {
  components: {
    SearchFilterBar,
    RecipeCard,
    Footer,
  },
  data() {
    return {
      menuOpen: false,
      bookmarkedRecipes: [], // Store bookmarked recipes
    };
  },
  methods: {
    toggleMenu() {
      this.menuOpen = !this.menuOpen;
    },
    goToPage(route) {
      this.$router.push(route);
    },
    createAccount() {
      this.$router.push('/create-account');
    },
    rateRecipe(rating) {
      this.$toast.success(`You rated this recipe ${rating} stars`);
    },
    commentOnRecipe() {
      this.$toast.info('Write your comment!');
    },
    bookmarkRecipe(recipe) {
      // Check if the recipe is already bookmarked
      const exists = this.bookmarkedRecipes.some((r) => r.id === recipe.id);
      if (!exists) {
        this.bookmarkedRecipes.push(recipe);
        this.$toast.success('Recipe added to bookmarks! Click to see.');
      } else {
        this.$toast.info('Recipe is already bookmarked.');
      }
    },
    goToPaymentPage() {
      this.$router.push('/payment');
    },
  },
};
</script>

<style scoped>
/* Menu button and dropdown menu responsiveness */
@media (max-width: 768px) {
  .menu-button {
    font-size: 1.25rem;
  }

  .dropdown-menu {
    width: 100%;
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
