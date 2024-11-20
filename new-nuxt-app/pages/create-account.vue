<template>
  <div class="min-h-screen bg-gradient-to-r from-teal-400 to-blue-500 flex items-center justify-center">
    <!-- Card Container -->
    <div class="w-full max-w-md bg-white p-8 rounded-lg shadow-xl transform transition-all duration-500 hover:scale-105">
      <!-- Title -->
      <h1 class="text-4xl font-semibold text-center text-gray-800 mb-6">Create Your Account</h1>
      
      <!-- Form -->
      <form @submit.prevent="submitForm">
        <div class="space-y-6">
          <!-- Full Name Input -->
          <div>
            <div class="relative">
              <input
                type="text"
                v-model="fullName"
                placeholder="Full Name"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
            </div>
          </div>

          <!-- Email Input -->
          <div>
            <div class="relative">
              <input
                type="email"
                v-model="email"
                placeholder="Email Address"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
            </div>
          </div>

          <!-- Password Input -->
          <div>
            <div class="relative">
              <input
                type="password"
                v-model="password"
                placeholder="Password"
                class="w-full p-4 pl-12 border rounded-lg shadow-sm focus:ring-2 focus:ring-teal-500 focus:outline-none transition"
              />
            </div>
          </div>

          <!-- Submit Button -->
          <div>
            <button
              type="submit"
              class="w-full bg-teal-500 text-white py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition duration-300 ease-in-out hover:bg-teal-600"
            >
              Create Account
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fullName: '',
      email: '',
      password: '',
    };
  },
  methods: {
    submitForm() {
      // Get initials from full name
      const initials = this.fullName.split(' ').map(word => word[0].toUpperCase()).join('');

      // Emit initials to parent component
      this.$emit('accountCreated', initials);

      // Redirect to homepage after successful account creation
      this.$router.push('/');
    },
  },
};
</script>

