<template>
  <div class="flex items-center justify-center h-screen bg-gray-100">
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

<script>
export default {
  data() {
    return {
      tempUser: null, // Temporary storage for user data
    };
  },
  created() {
    // Save the current user data temporarily
    const user = localStorage.getItem('user');
    if (user) {
      this.tempUser = JSON.parse(user);
    }
  },
  methods: {
    // Method to handle logout
    logout() {
      localStorage.removeItem('user'); // Clear user data from local storage
      this.$router.replace('/create-account'); // Use replace to navigate to login page
    },
    cancelLogout() {
      if (this.tempUser) {
        // Restore the user data if cancel is clicked
        localStorage.setItem('user', JSON.stringify(this.tempUser));
      }
      this.$router.push('/'); // Navigate back to the home page (default route)
    },
  },
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
