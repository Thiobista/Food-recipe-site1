<template>
  <div>
    <h2 class="text-3xl font-bold mb-6">Edit Recipe</h2>
    <form @submit.prevent="updateRecipe">
      <div class="mb-4">
        <label for="title" class="block text-lg font-semibold">Recipe Title</label>
        <input v-model="recipe.title" type="text" id="title" class="w-full p-2 border border-gray-300 rounded-lg" required />
      </div>
      
      <div class="mb-4">
        <label for="description" class="block text-lg font-semibold">Recipe Description</label>
        <textarea v-model="recipe.description" id="description" class="w-full p-2 border border-gray-300 rounded-lg" required></textarea>
      </div>
      
      <div class="mb-4">
        <label for="image" class="block text-lg font-semibold">Recipe Image URL</label>
        <input v-model="recipe.image" type="url" id="image" class="w-full p-2 border border-gray-300 rounded-lg" required />
      </div>

      <div v-if="loading" class="text-center mb-4">
        <span>Loading...</span> <!-- You can add a spinner here -->
      </div>

      <div v-if="error" class="text-center text-red-500 mb-4">
        <span>{{ error }}</span>
      </div>

      <div class="text-center">
        <button 
          type="submit" 
          :disabled="loading" 
          class="bg-blue-500 text-white py-2 px-6 rounded-lg hover:bg-blue-600 transition disabled:bg-gray-300">
          Save Changes
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// Define the reactive state
const recipe = ref({
  title: '',
  description: '',
  image: ''
})
const loading = ref(false)
const error = ref(null)

const route = useRoute()  // Access the current route (to get the recipe ID)
const router = useRouter() // Access the router to navigate after form submission

// Fetch the recipe by ID
const fetchRecipe = async (id) => {
  loading.value = true
  try {
    // Simulating fetching a recipe by ID (Replace with an API call)
    const fetchedRecipe = { id, title: 'Sample Recipe', description: 'This is a sample recipe.', image: '/images/salad.jpg' }
    recipe.value = fetchedRecipe
  } catch (err) {
    error.value = 'Failed to load the recipe'
  } finally {
    loading.value = false
  }
}

// Update the recipe
const updateRecipe = async () => {
  loading.value = true
  try {
    // Replace this with an actual API call to update the recipe
    console.log('Updated recipe:', recipe.value)
    router.push('/my-recipes') // Redirect back to My Recipes page
  } catch (err) {
    error.value = 'Failed to update the recipe'
  } finally {
    loading.value = false
  }
}

// Get the recipe ID from the route params and fetch the recipe
onMounted(() => {
  const recipeId = route.params.id
  fetchRecipe(recipeId)
})
</script>

<style scoped>
/* Style for the form */
form {
  max-width: 600px;
  margin: 0 auto;
}

input, textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

button {
  font-weight: 600;
}

button:disabled {
  background-color: #bbb;
}
</style>
