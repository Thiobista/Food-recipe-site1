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
  
  <script>
  export default {
    data() {
      return {
        recipe: {
          title: '',
          description: '',
        },
      };
    },
    async created() {
      const id = this.$route.params.id;
      // Fetch the recipe data using `id` (e.g., via API or mock data)
      this.recipe = await this.fetchRecipe(id);
    },
    methods: {
      async fetchRecipe(id) {
        // Simulated API call
        return { title: 'Sample Recipe', description: 'Sample Description' };
      },
      async updateRecipe() {
        // Simulate updating recipe
        console.log('Updated recipe:', this.recipe);
        this.$router.push('/'); // Navigate back to My Recipes after saving
      },
    },
  };
  </script>
  