<template>
  <div class="bg-gray-100 py-4">
    <div class="container mx-auto flex justify-between items-center">
      <!-- Search Input -->
      <input 
        v-model="searchQuery"
        type="text"
        placeholder="Search Recipes..."
        class="px-4 py-2 w-1/3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500"
      />
      <div class="flex space-x-4">
        <!-- Category Dropdown -->
        <select 
          v-model="selectedCategory"
          class="px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500"
        >
          <option value="">All Categories</option>
          <option value="Vegan">Vegan</option>
          <option value="Quick Meals">Quick Meals</option>
        </select>
      </div>
    </div>

    <!-- Display Filtered Recipes -->
    <div v-if="filteredRecipes.length > 0" class="mt-8 grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
      <div v-for="recipe in filteredRecipes" :key="recipe.id" class="recipe-card">
        <img :src="recipe.image" alt="Recipe Image" class="w-full h-48 object-cover rounded-lg"/>
        <div class="p-4">
          <h3 class="font-bold text-lg">{{ recipe.title }}</h3>
          <p class="text-sm text-gray-600">{{ recipe.category }}</p>
        </div>
      </div>
    </div>
    <p v-else class="mt-4 text-center text-gray-500">No recipes found</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      selectedCategory: '',
      recipes: [
        { id: 1, title: 'Vegan Salad', category: 'Vegan', image: '/path/to/salad.jpg' },
        { id: 2, title: 'Quick Pasta', category: 'Quick Meals', image: '/path/to/pasta.jpg' },
        { id: 3, title: 'Veggie Stir Fry', category: 'Vegan', image: '/path/to/stir-fry.jpg' },
        // Add more recipes as needed
      ]
    };
  },
  computed: {
    filteredRecipes() {
      return this.recipes.filter(recipe => {
        const matchesSearchQuery = recipe.title.toLowerCase().includes(this.searchQuery.toLowerCase());
        const matchesCategory = this.selectedCategory ? recipe.category === this.selectedCategory : true;
        return matchesSearchQuery && matchesCategory;
      });
    }
  }
};
</script>

<style scoped>
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
</style>
