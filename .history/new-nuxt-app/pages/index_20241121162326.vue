<template>
  <div>
    <!-- Navbar Section -->
    <header>
      <nav class="bg-yellow-500 p-4 flex justify-between items-center">
        <!-- Menu Button (Top Left) -->
        <div class="relative">
          <button
            @click="toggleMenu"
            class="text-white font-bold"
          >
            ☰ Menu
          </button>

          <!-- Dropdown Menu -->
          <div
            v-if="menuOpen"
            class="absolute bg-white shadow-lg rounded-lg mt-2 w-48 z-10"
          >
            <ul>
              <li class="p-2 hover:bg-yellow-100 flex items-center">
                <img src="/path/to/profile-image.jpg" alt="Profile" class="w-6 h-6 mr-2 rounded-full">
                Profile
              </li>
              <li class="p-2 hover:bg-yellow-100">Create Recipe</li>
              <li class="p-2 hover:bg-yellow-100">My Recipes</li>
              <li class="p-2 hover:bg-yellow-100">Bookmarks</li>
              <li class="p-2 hover:bg-yellow-100 text-red-500">Logout</li>
            </ul>
          </div>
        </div>

        <!-- Create Account Button (Top Right) -->
        <div class="auth-container">
  <button 
    class="w-full bg-blue-500 text-white py-2 rounded-lg"
    @click="createAccount"
  >
    Create Account
  </button>
</div>
      </nav>
    </header>

    <!-- Hero Section -->
    <HeroSection />

    <!-- Main Content -->
    <main class="container mx-auto py-8">
      <SearchFilterBar />

      <!-- Recipe Cards Grid -->
      <div class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-8">
        <RecipeCard
          v-for="n in 6"
          :key="n"
          image="/images/apple.jpg"
          title="Delicious Recipe"
          description="A quick preview of the recipe details."
        />
      </div>
    </main>

    <!-- Footer -->
    <Footer />
  </div>
</template>

<script>
import HeroSection from '~/components/HeroSection.vue'
import SearchFilterBar from '~/components/SearchFilterBar.vue'
import RecipeCard from '~/components/RecipeCard.vue'
import Footer from '~/components/Footer.vue'

export default {
  components: {
    HeroSection,
    SearchFilterBar,
    RecipeCard,
    Footer,
  },
  data() {
    return {
      menuOpen: false, // Toggle menu state
    };
  },
  methods: {
    toggleMenu() {
      this.menuOpen = !this.menuOpen; // Open/close the menu
    },
    createAccount() {
    this.$router.push('/create-account'); // Make sure this route matches an existing page.
  },
    goToAuthPage() {
      this.$router.push('/auth');  // Navigate to the authentication page
    }
  },
}
</script>

<style scoped>
/* Add styles for menu responsiveness and aesthetics */
</style><template> 
<div>
  <!-- Navbar Section -->
  <header>
    <nav class="bg-yellow-500 p-4 flex justify-between items-center">
      <!-- Menu Button (Top Left) -->
      <div class="relative">
        <button
          @click="toggleMenu"
          class="text-white font-bold"
        >
          ☰ Menu
        </button>

        <!-- Dropdown Menu -->
        <div
          v-if="menuOpen"
          class="absolute bg-white shadow-lg rounded-lg mt-2 w-48 z-10"
        >
          <ul>
            <li class="p-2 hover:bg-yellow-100 flex items-center">
              <img src="/path/to/profile-image.jpg" alt="Profile" class="w-6 h-6 mr-2 rounded-full">
              Profile
            </li>
            <li class="p-2 hover:bg-yellow-100">Create Recipe</li>
            <li class="p-2 hover:bg-yellow-100">My Recipes</li>
            <li class="p-2 hover:bg-yellow-100">Bookmarks</li>
            <li class="p-2 hover:bg-yellow-100 text-red-500">Logout</li>
          </ul>
        </div>
      </div><!-- Create Account Button (Top Right) -->
      <div class="auth-container">
        <button 
          class="w-full bg-blue-500 text-white py-2 rounded-lg"
          @click="createAccount"
        >
          Create Account
        </button>
      </div>
    </nav>
  </header>

  <!-- Hero Section -->
  <HeroSection />

  <!-- Main Content -->
  <main class="container mx-auto py-8">
    <SearchFilterBar />

    <!-- Recipe Cards Grid -->
    <div class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-8">
      <RecipeCard
        v-for="n in 6"
        :key="n"
        image="/images/enjera.jpg"
        title="Delicious Recipe"
        description="A quick preview of the recipe details."
      />
    </div>
  </main>

  <!-- Footer -->
  <Footer />
</div>
</template>

<script>
import HeroSection from '~/components/HeroSection.vue'
import SearchFilterBar from '~/components/SearchFilterBar.vue'
import RecipeCard from '~/components/RecipeCard.vue'
import Footer from '~/components/Footer.vue'

export default {
components: {
  HeroSection,
  SearchFilterBar,
  RecipeCard,
  Footer,
},
data() {
  return {
    menuOpen: false, // Toggle menu state
  };
},
methods: {
  toggleMenu() {
    this.menuOpen = !this.menuOpen; // Open/close the menu
  },
  createAccount() {
    // Change route programmatically
    this.$router.push({ path: '/create-account' });
  },
},
};
</script>

<style scoped>
/* Add styles for menu responsiveness and aesthetics */
</style>