<template>
  <div class="min-h-screen bg-gradient-to-br from-yellow-50 to-yellow-100 py-12 px-4">
    <!-- Header Section -->
    <div class="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
      <Header />
    </div>

    <!-- Content Section -->
    <div class="pt-20 px-4">
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
        <div v-for="recipe in recipes" :key="recipe.id" class="recipe-card bg-white p-4 rounded-lg shadow-md flex flex-col">
          <img :src="recipe.image || '/placeholder-image.jpg'" alt="Recipe" class="w-full h-48 object-cover rounded-lg" />
          <h3 class="text-xl font-semibold mt-4">{{ recipe.title }}</h3>
          <p class="text-sm text-gray-500 mt-2">{{ recipe.description }}</p>

          <!-- Edit, Delete, and Share buttons -->
          <div class="mt-auto flex justify-end space-x-2 pt-4">
            <button
              @click="editRecipe(recipe.id)"
              class="flex items-center justify-center bg-orange-500 text-white py-2 px-3 rounded-lg hover:bg-orange-600 transition shadow-md"
            >
              <font-awesome-icon icon="edit" />
            </button>
            <button
              @click="openDeleteModal(recipe.id)"
              class="flex items-center justify-center bg-red-500 text-white py-2 px-3 rounded-lg hover:bg-red-600 transition shadow-md"
            >
              <font-awesome-icon icon="trash" />
            </button>
            <button
              @click="shareRecipe(recipe)"
              class="flex items-center justify-center bg-orange-500 text-white py-2 px-3 rounded-lg hover:bg-orange-600 transition shadow-md"
            >
              <font-awesome-icon icon="share-alt" />
            </button>
          </div>
        </div>
      </div>

      <!-- Add New Recipe Button -->
      <div class="mt-8 text-center">
        <button
          @click="goToAddRecipe"
          class="bg-yellow-500 text-white py-2 px-6 rounded-full hover:bg-yellow-600 transition shadow-lg"
        >
          Add New Recipe
        </button>
      </div>
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
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

// Use the Nuxt router instance
const router = useRouter();

// Recipes and related states
const recipes = ref([]);
const showDeleteModal = ref(false);
const recipeToDelete = ref(null);
const showSuccessMessage = ref(false);
const message = ref('');

// GraphQL endpoint and headers
const GRAPHQL_ENDPOINT = 'https://your-hasura-instance/v1/graphql';
const AUTH_TOKEN = 'your-jwt-token'; // Replace with your authentication method

// Fetch recipes from Hasura GraphQL API
const fetchRecipes = async () => {
  const query = `
    query FetchRecipes {
      recipes {
        id
        title
        description
        image
      }
    }
  `;

  try {
    const response = await axios.post(
      GRAPHQL_ENDPOINT,
      { query },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${AUTH_TOKEN}`,
        },
      }
    );
    recipes.value = response.data.data.recipes;
  } catch (error) {
    console.error('Error fetching recipes:', error);
  }
};

// Navigation functions
const goToAddRecipe = () => {
  router.push('/create-recipe');
};
const editRecipe = (recipeId) => {
  router.push(`/recipes/edit/${recipeId}`);
};

// Modal and deletion functions
const openDeleteModal = (recipeId) => {
  recipeToDelete.value = recipeId;
  showDeleteModal.value = true;
};
const closeDeleteModal = () => {
  showDeleteModal.value = false;
  recipeToDelete.value = null;
};
const deleteRecipe = async () => {
  const mutation = `
    mutation DeleteRecipe($id: uuid!) {
      delete_recipes_by_pk(id: $id) {
        id
      }
    }
  `;
  try {
    await axios.post(
      GRAPHQL_ENDPOINT,
      {
        query: mutation,
        variables: { id: recipeToDelete.value },
      },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${AUTH_TOKEN}`,
        },
      }
    );
    recipes.value = recipes.value.filter((recipe) => recipe.id !== recipeToDelete.value);
    closeDeleteModal();
    message.value = 'Recipe deleted successfully';
    showSuccessMessage.value = true;
    setTimeout(() => (showSuccessMessage.value = false), 3000);
  } catch (error) {
    console.error('Error deleting recipe:', error);
  }
};

// Share recipe functionality
const shareRecipe = (recipe) => {
  const shareText = `${recipe.title} - ${recipe.description}`;
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
        setTimeout(() => (showSuccessMessage.value = false), 3000);
      })
      .catch((error) => console.error('Error sharing:', error));
  } else {
    const shareLink = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(shareUrl)}`;
    window.open(shareLink, '_blank');
  }
};


// Fetch recipes on page load
onMounted(fetchRecipes);



fetch("http://localhost:8081/create-recipe", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    title: "Recipe Title",
    description: "Recipe Description",
    ingredients: "Ingredients",
    steps: "Steps",
    time: "Time",
    category: "Category",
    userId: "UserID",
  }),
})
  .then((response) => response.json())
  .then((data) => console.log(data))
  .catch((error) => console.error("Error:", error));

</script>

<style scoped>
/* Styles omitted for brevity. Use the styles from your existing code */
</style>
