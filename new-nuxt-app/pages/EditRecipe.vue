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
  
        <div class="text-center">
          <button type="submit" class="bg-blue-500 text-white py-2 px-6 rounded-lg hover:bg-blue-600 transition">
            Save Changes
          </button>
        </div>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipe: {
          title: '',
          description: '',
          image: ''
        }
      };
    },
    created() {
      const recipeId = this.$route.params.id; // Get the recipe ID from the route
      this.fetchRecipe(recipeId); // Fetch the recipe details from the API
    },
    methods: {
      // Fetch the recipe by ID
      fetchRecipe(id) {
        // Replace with an API call to get the recipe by ID
        const fetchedRecipe = { id, title: 'Sample Recipe', description: 'This is a sample recipe.', image: '/images/sample.jpg' }; 
        this.recipe = fetchedRecipe;
      },
      
      // Update the recipe
      updateRecipe() {
        // Send a request to your API to update the recipe
        console.log('Updated recipe:', this.recipe);
        this.$router.push('/my-recipes'); // Redirect back to My Recipes page
      },
    },
  };
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
  </style>
  