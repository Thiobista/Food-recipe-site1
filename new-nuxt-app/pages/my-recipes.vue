<template>
  <div>
    <!-- My Recipes Page Heading -->
    <h2 class="text-3xl font-bold mb-6">My Recipes</h2>

    <!-- No Recipes Message -->
    <div v-if="recipes.length === 0" class="text-center text-gray-500">
      <p>You haven't added any recipes yet.</p>
    </div>

    <!-- Recipes Grid -->
    <div v-else class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-8">
      <!-- Loop through the user's recipes -->
      <div v-for="recipe in recipes" :key="recipe.id" class="recipe-card">
        <img :src="recipe.image" alt="Recipe" class="w-full h-48 object-cover rounded-lg" />
        <h3 class="text-xl font-semibold mt-4">{{ recipe.title }}</h3>
        <p class="text-sm text-gray-500 mt-2">{{ recipe.description }}</p>

        <!-- Edit, Delete, and Share buttons -->
        <div class="flex justify-between mt-4">
          <button
            @click="editRecipe(recipe.id)"
            class="bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 transition"
          >
            Edit
          </button>
          <button
            @click="confirmDeleteRecipe(recipe.id)"
            class="bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-600 transition"
          >
            Delete
          </button>
          <button
            @click="shareRecipe(recipe)"
            class="bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600 transition"
          >
            Share
          </button>
        </div>
      </div>
    </div>

    <!-- Add New Recipe Button -->
    <div class="mt-8 text-center">
      <button
        @click="goToAddRecipe"
        class="bg-yellow-500 text-white py-2 px-6 rounded-lg hover:bg-yellow-600 transition"
      >
        Add New Recipe
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      recipes: [
        { id: 1, title: 'Chocolate Cake', description: 'A delicious chocolate cake recipe.', image: '/images/chocolate-cake.jpg' },
        { id: 2, title: 'Apple Pie', description: 'A traditional apple pie recipe.', image: '/images/apple-pie.jpg' },
        // Add more sample recipes or fetch them from your backend
      ],
    };
  },
  methods: {
    goToAddRecipe() {
      this.$router.push('/create-recipe');
    },
    editRecipe(recipeId) {
      this.$router.push(`/recipes/edit/${recipeId}`);
    },
    confirmDeleteRecipe(recipeId) {
      if (confirm('Are you sure you want to delete this recipe?')) {
        this.deleteRecipe(recipeId);
      }
    },
    deleteRecipe(recipeId) {
      // Find the index of the recipe to delete
      const index = this.recipes.findIndex((recipe) => recipe.id === recipeId);

      if (index !== -1) {
        // Remove the recipe from the array
        this.recipes.splice(index, 1);
        this.$toast.success('Recipe deleted successfully');
      } else {
        this.$toast.error('Recipe not found');
      }
    },
    shareRecipe(recipe) {
      const shareText = `${recipe.title} - ${recipe.description}\nCheck it out!`;
      const shareUrl = window.location.href;
      const shareData = {
        title: recipe.title,
        text: shareText,
        url: shareUrl,
      };

      if (navigator.share) {
        navigator.share(shareData)
          .then(() => {
            this.$toast.success('Recipe shared successfully');
          })
          .catch((error) => {
            console.error('Error sharing:', error);
          });
      } else {
        const shareLink = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(shareUrl)}`;
        window.open(shareLink, '_blank');
      }
    },
  },
};
</script>

<style scoped>
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
button {
  font-weight: 600;
  text-transform: uppercase;
}
.recipe-card h3 {
  margin-top: 1rem;
}
</style>
