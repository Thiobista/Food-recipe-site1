export default defineNuxtConfig({
  vite: {
    server: {
      hmr: {
        protocol: 'ws', // Use 'wss' if using HTTPS
        host: 'localhost',
      },
    },
  },
  nitro: {
    logLevel: 'debug',
  },
  compatibilityDate: '2024-11-23',
  css: [
    "@/assets/css/main.css", // Main Tailwind CSS
    "@fortawesome/fontawesome-svg-core/styles.css", // Font Awesome styles
  ],
  postcss: {
    plugins: {
      tailwindcss: {}, // Tailwind configuration
      autoprefixer: {}, // Auto-prefixing CSS for browser compatibility
    },
  },
  modules: [
    "@nuxt/image-edge", // For Nuxt Image optimization
  ],
  build: {
    transpile: [
      "@apollo/client",
      "graphql",
      "@fortawesome/fontawesome-svg-core",
      "@fortawesome/free-solid-svg-icons",
      "@fortawesome/free-brands-svg-icons",
      "@fortawesome/vue-fontawesome",
    ], // Ensure Font Awesome is transpiled
  },
  runtimeConfig: {
    public: {
      apiUrl: process.env.API_URL || 'http://localhost:3000',
      graphqlEndpoint: "http://localhost:8080/v1/graphql", // Hasura GraphQL endpoint
    },
  },
  plugins: [
    '~/plugins/fontawesome.js', // Register the Font Awesome plugin
  ],
});
