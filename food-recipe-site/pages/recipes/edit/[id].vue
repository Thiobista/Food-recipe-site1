<template>
  <div>
    <h2 class="text-2xl font-bold">Edit Recipe</h2>
    <form @submit.prevent="updateRecipe">
      <div>
        <label for="title">Title:</label>
        <input v-model="recipe.title" type="text" id="title" required />
      </div>
      <div>
        <label for="description">Description:</label>
        <textarea v-model="recipe.description" id="description" required></textarea>
      </div>
      <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded">
        Save Changes
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// Reactive reference for recipe
const recipe = ref({
  title: '',
  description: '',
})

// Get route parameters
const route = useRoute()
const router = useRouter()

// Fetch recipe on mount
onMounted(async () => {
  const id = route.params.id
  recipe.value = await fetchRecipe(id)
})

// Simulated API call to fetch recipe by ID
async function fetchRecipe(id) {
  // Here you would make an API call
  return { title: 'Sample Recipe', description: 'Sample Description' }
}

// Method to update recipe
async function updateRecipe() {
  console.log('Updated recipe:', recipe.value)
  router.push('/') // Navigate back to My Recipes after saving
}
</script>
