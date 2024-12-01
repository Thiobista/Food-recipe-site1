<template>
  <div class="flex items-center justify-center h-screen bg-gray-100">
     <!-- Header Section -->
     <div class="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
      <Header />
    </div>
    <div class="bg-white p-8 rounded-lg shadow-md text-center w-80">
      <h2 class="text-2xl font-semibold text-gray-800 mb-4">Logout</h2>
      <p class="text-gray-600 mb-6">Are you sure you want to log out?</p>
      <button
        @click="logout"
        class="w-full bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-600 transition-all"
      >
        Logout
      </button>
      <button
        @click="cancelLogout"
        class="w-full mt-3 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg hover:bg-gray-300 transition-all"
      >
        Cancel
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

// Setup variables
const tempUser = ref(null);
const router = useRouter();

// onMounted lifecycle to handle initial setup
onMounted(() => {
  const user = localStorage.getItem("user");
  if (user) {
    tempUser.value = JSON.parse(user);
  }
});

// Logout method
const logout = () => {
  localStorage.removeItem("user"); // Clear user data from local storage
  router.replace("/create-account"); // Navigate to create account page
};

// Cancel logout method
const cancelLogout = () => {
  if (tempUser.value) {
    localStorage.setItem("user", JSON.stringify(tempUser.value));
  }
  router.push("/"); // Navigate back to the home page
};
</script>

<style scoped>
h2 {
  font-family: 'Inter', sans-serif;
}

p {
  font-family: 'Roboto', sans-serif;
}
</style>
