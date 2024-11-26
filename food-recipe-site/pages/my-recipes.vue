<template>
  <div>
    <!-- Success Message -->
    <div v-if="showSuccessMessage" class="fixed top-4 right-4 bg-green-500 text-white py-2 px-4 rounded-lg shadow-lg">
      <p>{{ message }}</p>
    </div>

    <!-- My Recipes Page Heading -->
    <h2 class="text-3xl font-bold mb-6 text-center">My Recipes</h2>

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
         edit
          </button>
          <button
            @click="openDeleteModal(recipe.id)"
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

    <!-- Delete Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-gray-900 bg-opacity-50 flex justify-center items-center z-50">
      <div class="bg-white p-6 rounded-lg shadow-xl max-w-sm w-full">
        <h3 class="text-lg font-semibold mb-4">Are you sure you want to delete this recipe?</h3>
        <div class="flex justify-between">
          <button @click="deleteRecipe" class="bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-600 transition">
            Yes, Delete
          </button>
          <button @click="closeDeleteModal" class="bg-gray-300 text-black py-2 px-4 rounded-lg hover:bg-gray-400 transition">
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

// Use the Nuxt router instance
const router = useRouter();

// Example recipe data (replace with actual fetch or state management)
const recipes = ref([
  { id: 1, title: 'Chocolate Cake', description: 'A delicious chocolate cake recipe.', image: '/images/chocolate.jpg' },
  { id: 2, title: 'Apple Pie', description: 'A traditional apple pie recipe.', image: '/images/apple.jpg' },
]);

// Delete modal visibility and selected recipe ID
const showDeleteModal = ref(false);
const recipeToDelete = ref(null);

// Success message visibility and content
const showSuccessMessage = ref(false);
const message = ref('');

// Navigation to Add New Recipe Page
const goToAddRecipe = () => {
  router.push('/create-recipe');
};

// Navigate to Edit Recipe Page
const editRecipe = (recipeId) => {
  router.push(`/recipes/edit/${recipeId}`);
};

// Open delete confirmation modal
const openDeleteModal = (recipeId) => {
  recipeToDelete.value = recipeId;
  showDeleteModal.value = true;
};

// Close delete confirmation modal
const closeDeleteModal = () => {
  showDeleteModal.value = false;
  recipeToDelete.value = null;
};

// Confirm and delete recipe
const deleteRecipe = () => {
  if (recipeToDelete.value !== null) {
    const index = recipes.value.findIndex((recipe) => recipe.id === recipeToDelete.value);
    if (index !== -1) {
      recipes.value.splice(index, 1); // Remove the recipe from the list
      closeDeleteModal(); // Close the modal after deletion
      // Show success message
      message.value = 'Recipe deleted successfully';
      showSuccessMessage.value = true;
      setTimeout(() => { showSuccessMessage.value = false; }, 3000); // Hide after 3 seconds
    } else {
      alert('Recipe not found');
    }
  }
};

// Share recipe functionality
const shareRecipe = (recipe) => {
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
        message.value = 'Recipe shared successfully';
        showSuccessMessage.value = true;
        setTimeout(() => { showSuccessMessage.value = false; }, 3000);
      })
      .catch((error) => console.error('Error sharing:', error));
  } else {
    const shareLink = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(shareUrl)}`;
    window.open(shareLink, '_blank');
  }
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

/* Modal Styles */
.fixed {
  position: fixed;
}
.bg-opacity-50 {
  background-color: rgba(0, 0, 0, 0.5);
}
.bg-white {
  background-color: #ffffff;
}
.p-6 {
  padding: 1.5rem;
}
.rounded-lg {
  border-radius: 8px;
}
.shadow-xl {
  box-shadow: 0 10px 15px rgba(0, 0, 0, 0.1);
}

/* Success message styles */
.fixed.top-4.right-4 {
  z-index: 1000;
}
</style>
