<template>
    <div>
      <h1 class="text-2xl font-bold">Create a Recipe</h1>
      <form @submit.prevent="submitForm">
        <VeeValidateForm>
          <!-- Recipe Title -->
          <Field name="title" as="input" v-model="form.title" placeholder="Title" rules="required" />
          <ErrorMessage name="title" />
  
          <!-- Add Ingredients -->
          <div>
            <h3>Ingredients</h3>
            <div v-for="(ingredient, index) in form.ingredients" :key="index">
              <input v-model="form.ingredients[index]" placeholder="Ingredient" />
              <button @click="removeIngredient(index)">Remove</button>
            </div>
            <button @click="addIngredient">Add Ingredient</button>
          </div>
  
          <!-- Submit -->
          <button type="submit">Create Recipe</button>
        </VeeValidateForm>
      </form>
    </div>
  </template>
  
  <script setup>
  import { useForm, Field, ErrorMessage, VeeValidateForm } from "vee-validate";
  
  const form = reactive({
    title: "",
    ingredients: [""],
  });
  
  const addIngredient = () => form.ingredients.push("");
  const removeIngredient = (index) => form.ingredients.splice(index, 1);
  
  const submitForm = () => {
    console.log(form);
  };
  </script>
  