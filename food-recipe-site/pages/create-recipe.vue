<template>
  <div class="min-h-screen bg-gradient-to-br from-yellow-50 to-yellow-100 py-12 px-4">
      <!-- Header Section -->
      <div class="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
      <Header />
    </div>
          <!-- Content Section -->
    <div class="pt-20 px-4"> <!-- Added padding to push content below the fixed header -->
    </div>
    <div class="max-w-3xl mx-auto bg-white rounded-lg shadow-lg p-8">
      <h1 class="text-4xl font-bold text-gray-800 text-center mb-8">Create a New Recipe</h1>

      <form @submit.prevent="submitRecipe" class="space-y-6">
        <!-- Recipe Title -->
        <div>
          <label for="recipe-title" class="block text-lg font-medium text-gray-700">Recipe Title</label>
          <input
            v-model="recipe.title"
            type="text"
            id="recipe-title"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            placeholder="Enter the recipe title"
            required
          />
        </div>

        <!-- Recipe Description -->
        <div>
          <label for="recipe-description" class="block text-lg font-medium text-gray-700">Description</label>
          <textarea
            v-model="recipe.description"
            id="recipe-description"
            rows="4"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            placeholder="Enter a brief description of the recipe"
            required
          ></textarea>
        </div>

        <!-- Ingredients -->
        <div>
          <label for="recipe-ingredients" class="block text-lg font-medium text-gray-700">Ingredients</label>
          <textarea
            v-model="recipe.ingredients"
            id="recipe-ingredients"
            rows="4"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            placeholder="Enter the ingredients"
            required
          ></textarea>
        </div>

        <!-- Steps -->
        <div>
          <label for="recipe-steps" class="block text-lg font-medium text-gray-700">Steps</label>
          <textarea
            v-model="recipe.steps"
            id="recipe-steps"
            rows="6"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            placeholder="Enter the cooking steps"
            required
          ></textarea>
        </div>

        <!-- Time to Prepare -->
        <div>
          <label for="recipe-time" class="block text-lg font-medium text-gray-700">Time to Prepare</label>
          <input
            v-model="recipe.time"
            type="text"
            id="recipe-time"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            placeholder="Enter the preparation time (e.g., 30 minutes)"
            required
          />
        </div>

        <!-- Food Category -->
        <div>
          <label for="recipe-category" class="block text-lg font-medium text-gray-700">Food Category</label>
          <select
            v-model="recipe.category"
            id="recipe-category"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
            required
          >
            <option value="">Select a category</option>
            <option value="Breakfast">Breakfast</option>
            <option value="Lunch">Lunch</option>
            <option value="Dinner">Dinner</option>
            <option value="Snack">Snack</option>
            <option value="Dessert">Dessert</option>
          </select>
        </div>

        <!-- Image Upload -->
        <div>
          <label for="recipe-image" class="block text-lg font-medium text-gray-700">Recipe Image</label>
          <input
            type="file"
            id="recipe-image"
            @change="handleImageUpload"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg shadow-sm focus:ring-yellow-400 focus:border-yellow-400"
          />
        </div>

        <!-- Submit Button -->
        <div class="text-center">
          <button
            type="submit"
            class="bg-yellow-500 text-white px-8 py-3 rounded-lg shadow-md hover:bg-yellow-600 transform hover:scale-105 transition duration-300"
          >
            Submit Recipe
          </button>
        </div>
      </form>
    </div>
  </div>

</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

const recipe = ref({
    title: '',
    description: '',
    ingredients: '',
    steps: '',
    time: '',
    category: '',
    image: null,
});

const router = useRouter();

const handleImageUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
            recipe.value.image = e.target.result; // Save the base64 image
        };
        reader.readAsDataURL(file);
    }
};

const submitRecipe = async () => {
    try {
        // Make the POST request
        const response = await axios.post('http://localhost:8081/create-recipe', {
            title: recipe.value.title,
            description: recipe.value.description,
            ingredients: recipe.value.ingredients,
            steps: recipe.value.steps,
            time: recipe.value.time,
            category: recipe.value.category,
            image: recipe.value.image,
            userId: 'exampleUserId123', // Replace with the actual logged-in user's ID
        });

        // Redirect or show success message
        console.log(response.data.message);
        router.push('/my-recipes'); // Navigate to the homepage or another page
    } catch (error) {
        console.error('Error submitting recipe:', error.response?.data || error.message);
        alert('Failed to submit the recipe. Please try again.');
    }
};

</script>



<style scoped>
body {
  font-family: 'Inter', sans-serif;
}

form input:focus,
form textarea:focus,
form select:focus {
  outline: none;
}

form {
  transition: transform 0.3s ease-in-out;
}

form:hover {
  transform: scale(1.02);
}
</style>
