<template>
    <div class="container mx-auto py-10 px-4">
      <Header />
      <!-- Page Title -->
      <h1 class="text-4xl font-bold text-yellow-500 text-center">Our Blog</h1>
      <p class="text-gray-700 mt-4 text-center">
        Stay updated with the latest news and cooking tips!
      </p>
  
      <!-- Blog Posts -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-10">
        <div
          v-for="(post, index) in posts"
          :key="index"
          class="bg-white rounded-lg shadow-md hover:shadow-lg transition"
        >
          <img
            :src="post.image"
            :alt="post.title"
            class="w-full h-40 object-cover rounded-t-lg"
          />
          <div class="p-4">
            <h2 class="text-lg font-semibold text-gray-800">{{ post.title }}</h2>
            <p class="text-sm text-gray-600 mt-2">{{ post.summary }}</p>
            <div class="mt-4 text-right">
              <button
                @click="readMore(post.id)"
                class="px-4 py-2 bg-yellow-500 text-white text-sm rounded hover:bg-yellow-600 transition"
              >
                Read More
              </button>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Pagination -->
      <div class="mt-8 flex justify-center">
        <button
          class="px-4 py-2 mx-1 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition"
          @click="prevPage"
          :disabled="currentPage === 1"
        >
          Previous
        </button>
        <span class="px-4 py-2 text-gray-700">{{ currentPage }}</span>
        <button
          class="px-4 py-2 mx-1 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition"
          @click="nextPage"
          :disabled="currentPage === totalPages"
        >
          Next
        </button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Blog",
    data() {
      return {
        posts: [
          {
            id: 1,
            title: "Blog Post Title 1",
            summary: "Summary of the blog post. Click to read more...",
            image: "https://via.placeholder.com/300x200",
          },
          {
            id: 2,
            title: "Blog Post Title 2",
            summary: "Summary of the blog post. Click to read more...",
            image: "https://via.placeholder.com/300x200",
          },
          {
            id: 3,
            title: "Blog Post Title 3",
            summary: "Summary of the blog post. Click to read more...",
            image: "https://via.placeholder.com/300x200",
          },
        ],
        currentPage: 1,
        totalPages: 3,
      };
    },
    methods: {
      readMore(postId) {
        // Navigate to the detailed blog post page
        this.$router.push(`/blog/${postId}`);
      },
      nextPage() {
        if (this.currentPage < this.totalPages) {
          this.currentPage++;
          // Fetch new posts for the current page (replace this with API call logic)
        }
      },
      prevPage() {
        if (this.currentPage > 1) {
          this.currentPage--;
          // Fetch new posts for the current page (replace this with API call logic)
        }
      },
    },
  };
  </script>
  
  <style scoped>
  button {
    outline: none;
  }
  </style>
  