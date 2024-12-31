<template>
  <div class="max-w-2xl mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Create a New Recipe</h1>

    <form @submit.prevent="submitRecipe">
      <div class="mb-4">
        <label for="title" class="block font-medium mb-2">Recipe Title</label>
        <input
          id="title"
          v-model="form.title"
          type="text"
          placeholder="Enter recipe title"
          class="border border-gray-300 rounded-md p-2 w-full"
          required
        />
      </div>

      <div class="mb-4">
        <label for="description" class="block font-medium mb-2">Description</label>
        <textarea
          id="description"
          v-model="form.description"
          placeholder="Enter recipe description"
          class="border border-gray-300 rounded-md p-2 w-full"
          rows="3"
        ></textarea>
      </div>

      <div class="mb-4">
        <label for="time" class="block font-medium mb-2">Time to Prepare (minutes)</label>
        <input
          id="time"
          v-model.number="form.time_to_prepare"
          type="number"
          min="1"
          placeholder="Enter time in minutes"
          class="border border-gray-300 rounded-md p-2 w-full"
          required
        />
      </div>

      <div class="mb-4">
        <label for="image" class="block font-medium mb-2">Featured Image</label>
        <input
          id="image"
          type="file"
          @change="handleImageUpload"
          class="border border-gray-300 rounded-md p-2 w-full"
          accept="image/*"
        />
        <div v-if="form.featured_image_url" class="mt-2">
          <img :src="form.featured_image_url" alt="Preview" class="max-h-40" />
        </div>
      </div>

      <div class="mb-4">
        <label for="category" class="block font-medium mb-2">Category</label>
        <select
          id="category"
          v-model="form.category_id"
          class="border border-gray-300 rounded-md p-2 w-full"
          required
        >
          <option value="" disabled>Select a category</option>
          <option v-for="category in categories" :key="category.id" :value="category.id">
            {{ category.name }}
          </option>
        </select>
      </div>

      <div class="mb-4">
        <label for="steps" class="block font-medium mb-2">Steps</label>
        <div v-for="(step, index) in form.steps" :key="index" class="flex items-center mb-2">
          <input
            v-model="form.steps[index].step_description"
            type="text"
            placeholder="Enter step description"
            class="border border-gray-300 rounded-md p-2 flex-1"
            required
          />
          <button type="button" @click="removeStep(index)" class="ml-2 text-red-500">
            Remove
          </button>
        </div>
        <button type="button" @click="addStep" class="text-blue-500">
          + Add Step
        </button>
      </div>

      <div class="mb-4">
        <label for="ingredients" class="block font-medium mb-2">Ingredients</label>
        <div v-for="(ingredient, index) in form.ingredients" :key="index" class="flex items-center mb-2">
          <input
            v-model="form.ingredients[index].ingredient_name"
            type="text"
            placeholder="Enter ingredient"
            class="border border-gray-300 rounded-md p-2 flex-1"
            required
          />
          <button type="button" @click="removeIngredient(index)" class="ml-2 text-red-500">
            Remove
          </button>
        </div>
        <button type="button" @click="addIngredient" class="text-blue-500">
          + Add Ingredient
        </button>
      </div>

      <button
        type="submit"
        class="bg-blue-500 text-white font-medium rounded-md p-2 w-full"
      >
        Submit Recipe
      </button>
    </form>
  </div>
</template>

<script>
import gql from "graphql-tag";
import { useMutation } from "@vue/apollo-composable";

const CREATE_RECIPE = gql`
  mutation CreateFullRecipe($input: recipes_insert_input!) {
    insert_recipes_one(object: $input) {
      id
      steps {
        step_number
        step_description
      }
      ingredients {
        ingredient_name
        quantity
      }
    }
  }
`;

export default {
  data() {
    return {
      form: {
        title: "",
        description: "",
        time_to_prepare: null,
        featured_image_url: "",
        category_id: "",
        steps: [{ step_number: 1, step_description: "" }],
        ingredients: [{ ingredient_name: "", quantity: "" }],
      },
      categories: [
        { id: "1", name: "Breakfast" },
        { id: "2", name: "Lunch" },
        { id: "3", name: "Dinner" },
      ],
    };
  },
  setup() {
    const { mutate: createRecipe } = useMutation(CREATE_RECIPE);
    return { createRecipe };
  },
  methods: {
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.form.featured_image_url = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    addStep() {
      const nextStep = this.form.steps.length + 1;
      this.form.steps.push({ step_number: nextStep, step_description: "" });
    },
    removeStep(index) {
      this.form.steps.splice(index, 1);
    },
    addIngredient() {
      this.form.ingredients.push({ ingredient_name: "", quantity: "" });
    },
    removeIngredient(index) {
      this.form.ingredients.splice(index, 1);
    },
    async submitRecipe() {
  const token = localStorage.getItem("token"); // Retrieve stored token
  if (!token) {
    alert("Authorization token not found");
    return;
  }

  // Prepare recipe data
  const recipeData = {
    title: this.form.title,
    description: this.form.description,
    time_to_prepare: this.form.time_to_prepare,
    featured_image_url: this.form.featured_image_url,
    category_id: this.form.category_id,
    user_id: "a55fe52d-6484-4482-824f-66f6fb659d30", // Use the actual user ID here
    steps: this.form.steps,
    ingredients: this.form.ingredients,
  };

  try {
      const response = await axios.post('http://localhost:8081/api/recipes', this.form, {
        headers: {
          'Authorization': `Bearer ${this.token}` // Pass the JWT token
        }
      });

    const result = await response.json();

    if (response.ok) {
      console.log("Recipe submitted successfully:", result);
      alert("Recipe submitted successfully!");
    } else {
      throw new Error(result.message || "Something went wrong.");
    }
  } catch (error) {
    console.error("Error creating recipe:", error);
    alert(`Error: ${error.message}`);
  }
}


  },
};
</script>
