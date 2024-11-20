<template>
  <div>
    <!-- Bookmarked Recipes Section -->
    <h2 class="text-2xl font-bold mb-6">Bookmarked Recipes</h2>

    <div v-if="bookmarkedRecipes.length === 0" class="text-center text-gray-500">
      <p>No recipes bookmarked yet.</p>
    </div>

    <div v-else class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-8">
      <!-- Displaying Bookmarked Recipes -->
      <div 
        v-for="recipe in bookmarkedRecipes" 
        :key="recipe.id" 
        class="recipe-card"
      >
        <img 
          :src="recipe.image" 
          :alt="`Image of ${recipe.title}`" 
          class="w-full h-48 object-cover rounded-lg"
        />
        <h3 class="text-xl font-semibold mt-4">{{ recipe.title }}</h3>
        <p class="text-sm text-gray-500 mt-2">{{ recipe.description }}</p>
        <button 
          @click="removeBookmark(recipe.id)" 
          class="mt-4 bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-600 transition"
        >
          Remove Bookmark
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      bookmarkedRecipes: [],
    };
  },
  mounted() {
    this.loadBookmarkedRecipes();
  },
  methods: {
    // Load bookmarked recipes from an API or local storage
    loadBookmarkedRecipes() {
      // Example static data; replace with dynamic fetch logic.
      this.bookmarkedRecipes = [
        { id: 1, title: 'Chocolate Cake', description: 'A delicious chocolate cake recipe.', image: '/images/chocolate-cake.jpg' },
        { id: 2, title: 'Apple Pie', description: 'A traditional apple pie recipe.', image: '/images/apple-pie.jpg' },
      ];
    },
    // Method to remove a recipe from bookmarks
    removeBookmark(recipeId) {
      this.bookmarkedRecipes = this.bookmarkedRecipes.filter(recipe => recipe.id !== recipeId);
    },
  },
};
</script>

<style scoped>
/* Styling for the recipe card */
.recipe-card {
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
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
}

button {
  font-weight: 600;
  text-transform: uppercase;
}
</style>
