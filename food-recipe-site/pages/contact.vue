<template>
  <div class="min-h-screen bg-gray-100 p-6">
    <Header />
    <h1 class="text-3xl font-bold text-center text-yellow-500">Contact Us</h1>
    <form @submit.prevent="submitForm" class="max-w-md mx-auto mt-6 bg-white p-4 shadow-md rounded-lg">
      <div class="mb-4">
        <label for="name" class="block text-gray-700">Name:</label>
        <input v-model="form.name" type="text" id="name" class="w-full border rounded p-2" required />
      </div>
      <div class="mb-4">
        <label for="email" class="block text-gray-700">Email:</label>
        <input v-model="form.email" type="email" id="email" class="w-full border rounded p-2" required />
      </div>
      <div class="mb-4">
        <label for="message" class="block text-gray-700">Message:</label>
        <textarea v-model="form.message" id="message" class="w-full border rounded p-2" required></textarea>
      </div>
      <button
        type="submit"
        class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600 transition"
      >
        Submit
      </button>
    </form>
  </div>
</template>

<script>
import { useToast } from "vue-toastification";
import { useRouter } from "vue-router";

export default {
  name: "ContactUs",
  data() {
    return {
      form: {
        name: "",
        email: "",
        message: "",
      },
    };
  },
  setup() {
    const toast = useToast();
    const router = useRouter();

    const submitForm = () => {
      // Display toast message
      toast.success("Your message has been submitted successfully!", {
        timeout: 3000,
        position: "top-right",
      });

      // Redirect to home page after a delay
      setTimeout(() => {
        router.push("/");
      }, 3000); // Adjust the delay as needed
    };

    return { submitForm };
  },
};
</script>

<style scoped>
button {
  outline: none;
}
</style>
