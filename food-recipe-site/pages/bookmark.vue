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
  
          <!-- Form Validation for Removing Bookmark -->
          <Form @submit="handleRemoveBookmark(recipe.id)">
            <button
              type="submit"
              class="mt-4 bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-600 transition"
            >
              Remove Bookmark
            </button>
          </Form>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import { Form } from 'vee-validate';
  import { useQuery, useMutation } from '@vue/apollo-composable';
  
  // Apollo Query for loading bookmarked recipes
  const GET_BOOKMARKED_RECIPES = `
    query GetBookmarkedRecipes {
      bookmarkedRecipes {
        id
        title
        description
        image
      }
    }
  `;
  
  // Apollo Mutation for removing a bookmark
  const REMOVE_BOOKMARK = `
    mutation RemoveBookmark($id: ID!) {
      removeBookmark(id: $id) {
        success
        message
      }
    }
  `;
  
  // Reactive state for bookmarked recipes
  const bookmarkedRecipes = ref([]);
  
  // Load bookmarked recipes with Apollo query
  const { result: recipesResult, loading, error } = useQuery(GET_BOOKMARKED_RECIPES);
  
  // Fetch data on component mount
  //  error, loading } = useQuery(GET_BOOKMARKED_RECIPES);
  
  // Populate bookmarkedRecipes when data is loaded
  onMounted(() => {
    recipesResult.onResult(({ data }) => {
      if (data && data.bookmarkedRecipes) {
        bookmarkedRecipes.value = data.bookmarkedRecipes;
      }
    });
  });
  
  // Mutation to handle removing a bookmark
  const { mutate: removeBookmark, onDone } = useMutation(REMOVE_BOOKMARK);
  
  // Handle bookmark removal
  const handleRemoveBookmark = (id) => {
    removeBookmark({ id });
    onDone(({ data }) => {
      if (data.removeBookmark.success) {
        // Update UI after successful removal
        bookmarkedRecipes.value = bookmarkedRecipes.value.filter(
          (recipe) => recipe.id !== id
        );
        alert("Bookmark removed successfully.");
      } else {
        alert("Failed to remove bookmark. Please try again.");
      }
    });
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
  