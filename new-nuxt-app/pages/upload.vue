<template>
    <section class="container mx-auto py-12">
      <h2 class="text-2xl font-bold mb-6">Upload a Recipe</h2>
      <RecipeUploadForm />
    </section>
  </template>
  
  <script>
  import RecipeUploadForm from '~/components/RecipeUploadForm.vue'
  
  export default {
    components: {
      RecipeUploadForm,
    },
  }
  </script>

