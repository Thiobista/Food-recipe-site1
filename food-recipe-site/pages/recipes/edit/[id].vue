<template>
  <div class="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-lg">
    <Header />
    <h2 class="text-3xl font-semibold text-center text-gray-800 mb-6">Edit Recipe</h2>
    <form @submit.prevent="updateRecipe" class="space-y-6">
      <!-- Recipe Title -->
      <div>
        <label for="title" class="block text-lg font-medium text-gray-700">Recipe Title:</label>
        <input 
          v-model="recipe.title" 
          type="text" 
          id="title" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          placeholder="Enter the recipe title" 
          required 
        />
      </div>

      <!-- Description -->
      <div>
        <label for="description" class="block text-lg font-medium text-gray-700">Description:</label>
        <textarea 
          v-model="recipe.description" 
          id="description" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          placeholder="Enter a brief description of the recipe" 
          required 
        ></textarea>
      </div>

      <!-- Ingredients -->
      <div>
        <label for="ingredients" class="block text-lg font-medium text-gray-700">Ingredients:</label>
        <textarea 
          v-model="recipe.ingredients" 
          id="ingredients" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          placeholder="Enter the ingredients" 
          required 
        ></textarea>
      </div>

      <!-- Steps -->
      <div>
        <label for="steps" class="block text-lg font-medium text-gray-700">Steps:</label>
        <textarea 
          v-model="recipe.steps" 
          id="steps" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          placeholder="Enter the cooking steps" 
          required 
        ></textarea>
      </div>

      <!-- Time to Prepare -->
      <div>
        <label for="time" class="block text-lg font-medium text-gray-700">Time to Prepare:</label>
        <input 
          v-model="recipe.time" 
          type="text" 
          id="time" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          placeholder="Enter the preparation time (e.g., 30 minutes)" 
          required 
        />
      </div>

      <!-- Food Category -->
      <div>
        <label for="category" class="block text-lg font-medium text-gray-700">Food Category:</label>
        <select 
          v-model="recipe.category" 
          id="category" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
          required 
        >
          <option disabled value="">Select a category</option>
          <option value="Appetizer">Appetizer</option>
          <option value="Main Course">Main Course</option>
          <option value="Dessert">Dessert</option>
          <option value="Beverage">Beverage</option>
        </select>
      </div>

      <!-- Recipe Image -->
      <div>
        <label for="image" class="block text-lg font-medium text-gray-700">Recipe Image:</label>
        <input 
          type="file" 
          id="image" 
          @change="handleFileUpload" 
          class="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" 
        />
        <p v-if="recipe.imageName" class="text-sm mt-2 text-gray-500">Selected File: {{ recipe.imageName }}</p>
      </div>

      <!-- Submit Button -->
      <div class="flex justify-center">
        <button 
          type="submit" 
          class="bg-yellow-600 text-white py-3 px-6 rounded-lg text-lg font-semibold hover:bg-blue-700 transition duration-300"
        >
          Save Changes
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// Reactive reference for recipe
const recipe = ref({
  title: '',
  description: '',
  ingredients: '',
  steps: '',
  time: '',
  category: '',
  imageName: '', // To display the selected image name
});

// Get route parameters
const route = useRoute();
const router = useRouter();

// Fetch recipe on mount
onMounted(async () => {
  const id = route.params.id;
  recipe.value = await fetchRecipe(id);
});

// Simulated API call to fetch recipe by ID
async function fetchRecipe(id) {
  // Replace with an actual API call
  return {
    title: 'Sample Recipe',
    description: 'A delicious sample recipe.',
    ingredients: 'Sample ingredients list.',
    steps: '1. Sample step 1\n2. Sample step 2',
    time: '30 minutes',
    category: 'Main Course',
  };
}

// Handle file upload
function handleFileUpload(event) {
  const file = event.target.files[0];
  if (file) {
    recipe.value.imageName = file.name;
    // You can also handle file uploading logic here
  }
}

// Method to update recipe
async function updateRecipe() {
  console.log('Updated recipe:', recipe.value);
  router.push('/my-recipes'); // Navigate back to the home or My Recipes page after saving
}
</script>

<style scoped>
/* Add any styles if necessary */
</style>
