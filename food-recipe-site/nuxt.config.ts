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
  css: ["@/assets/css/main.css"],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: ["@nuxt/image-edge"],
  build: {
    transpile: ["@apollo/client", "graphql"],
  },
  runtimeConfig: {
    public: {
      apiUrl: process.env.API_URL || 'http://localhost:3000',
      graphqlEndpoint: "http://localhost:8080/v1/graphql", // Hasura GraphQL endpoint
    },
  },
});
