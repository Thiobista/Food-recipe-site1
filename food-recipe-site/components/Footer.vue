<template>
  <footer class="bg-gray-900 text-white py-10">
    <div class="container mx-auto px-6 lg:px-20">
      <!-- Top Section -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 text-center md:text-left">
        <!-- About Us -->
        <div>
          <h2 class="text-xl font-bold text-yellow-500 mb-4">About Us</h2>
          <p class="text-sm text-gray-400 leading-relaxed">
            RecipeSite is your go-to platform for discovering, creating, and sharing amazing recipes. Join our community and ignite your culinary passion!
          </p>
        </div>

        <!-- Quick Links -->
        <div>
          <h2 class="text-xl font-bold text-yellow-500 mb-4">Quick Links</h2>
          <ul class="space-y-2">
            <li><a href="#" class="hover:text-yellow-400 transition-colors">Home</a></li>
            <li><a href="#" class="hover:text-yellow-400 transition-colors">Recipes</a></li>
            <li><a href="#" class="hover:text-yellow-400 transition-colors">Blog</a></li>
            <li><a href="#" class="hover:text-yellow-400 transition-colors">Contact Us</a></li>
            <li><a href="#" class="hover:text-yellow-400 transition-colors">Privacy Policy</a></li>
          </ul>
        </div>

        <!-- Social Media -->
        <div>
          <h2 class="text-xl font-bold text-yellow-500 mb-4">Follow Us</h2>
          <div class="flex justify-center md:justify-start space-x-6">
            <a href="#" class="text-gray-400 hover:text-yellow-400 transition transform hover:scale-110">
              <font-awesome-icon :icon="['fab', 'facebook']" class="text-2xl" />
            </a>
            <a href="#" class="text-gray-400 hover:text-yellow-400 transition transform hover:scale-110">
              <font-awesome-icon :icon="['fab', 'instagram']" class="text-2xl" />
            </a>
            <a href="#" class="text-gray-400 hover:text-yellow-400 transition transform hover:scale-110">
              <font-awesome-icon :icon="['fab', 'twitter']" class="text-2xl" />
            </a>
            <a href="#" class="text-gray-400 hover:text-yellow-400 transition transform hover:scale-110">
              <font-awesome-icon :icon="['fab', 'youtube']" class="text-2xl" />
            </a>
          </div>
        </div>
      </div>

      <!-- Divider -->
      <div class="border-t border-gray-700 my-8"></div>

      <!-- Bottom Section -->
      <div class="flex flex-col md:flex-row justify-between items-center text-sm text-gray-900">
        <p class="mb-4 md:mb-0">&copy; 2024 RecipeSite. All rights reserved.</p>
        <p>
          Designed with ❤️ by
          <a href="#" class="hover:text-yellow-400 transition-colors font-semibold">Thiobista</a>
        </p>
      </div>
    </div>
  </footer>
</template>

<style scoped>
footer {
  background: linear-gradient(to bottom, #494949, #807e7e);
}
</style>
