<template>
  <div class="relative min-h-screen bg-cover bg-center flex items-center justify-center">
    <div class="absolute inset-0 bg-cover bg-center blur-sm" style="background-image: url('images/chief.jpg');"></div>

    <div class="relative w-full max-w-md bg-white p-8 rounded-lg shadow-xl transform transition-all duration-500 hover:scale-105">
      <h1 class="text-4xl font-semibold text-center text-gray-800 mb-6">Create Your Account</h1>

      <form @submit.prevent="submitForm">
        <div class="space-y-6">
          <div>
            <input
              type="text"
              v-model="fullName"
              placeholder="Full Name"
              :class="{'border-red-500': errors.fullName}"
              class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
            />
            <span v-if="errors.fullName" class="text-red-500 text-sm">{{ errors.fullName }}</span>
          </div>
          <div>
            <input
              type="email"
              v-model="email"
              placeholder="Email Address"
              :class="{'border-red-500': errors.email}"
              class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
            />
            <span v-if="errors.email" class="text-red-500 text-sm">{{ errors.email }}</span>
          </div>
          <div>
            <input
              type="password"
              v-model="password"
              placeholder="Password"
              :class="{'border-red-500': errors.password}"
              class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
            />
            <span v-if="errors.password" class="text-red-500 text-sm">{{ errors.password }}</span>
          </div>
          <div>
            <button
              :disabled="loading"
              type="submit"
              class="w-full bg-teal-500 text-white py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-teal-600 disabled:opacity-50"
            >
              {{ loading ? 'Creating Account...' : 'Create Account' }}
            </button>
          </div>
        </div>
      </form>

      <div class="mt-4 text-center">
        <button
          @click="continueWithoutAccount"
          class="w-full bg-blue-100 text-blue-700 py-2 px-4 rounded-lg shadow-md hover:bg-blue-200 transform hover:scale-105 transition duration-300"
        >
          Continue Without Account
        </button>
      </div>

      <div class="mt-4 text-center">
        <button
          @click="goToLogin"
          class="w-full bg-gray-100 text-gray-700 py-2 px-4 rounded-lg shadow-md hover:bg-gray-200 transform hover:scale-105 transition duration-300"
        >
          Already have an account? Login
        </button>
      </div>
    </div>
  </div>
</template>


<script setup>
import { ref } from 'vue';
import { useRouter } from '#app';
import axios from 'axios';
import { useToast } from 'vue-toastification';

const router = useRouter();
const toast = useToast();
const fullName = ref('');
const email = ref('');
const password = ref('');
const errors = ref({});
const loading = ref(false);

const submitForm = async () => {
  errors.value = {};
  if (!fullName.value) errors.value.fullName = 'Full name is required';
  if (!email.value) errors.value.email = 'Email is required';
  else if (!/\S+@\S+\.\S+/.test(email.value)) errors.value.email = 'Email is invalid';
  if (!password.value) errors.value.password = 'Password is required';

  if (Object.keys(errors.value).length === 0) {
    try {
      loading.value = true;
      const response = await axios.post('http://localhost:8081/signup', {
        username: fullName.value,
        email: email.value,
        password: password.value,
      });

      if (response.data.userId) {
        toast.success('Account created successfully!');
        router.push('/login');
      }
    } catch (error) {
      console.error('Error during signup:', error);
      toast.error('An unexpected error occurred. Please try again later.');
    } finally {
      loading.value = false;
    }
  }
};

const continueWithoutAccount = () => {
  router.push('/');
};

const goToLogin = () => {
  router.push('/login');
};
</script>
