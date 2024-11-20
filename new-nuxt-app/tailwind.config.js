/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{vue,js,ts}',
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.{vue,js,ts}',
    './plugins/**/*.{js,ts}',
    './app.vue', // Include Nuxt 3's main entry file
  ],
  theme: {
    extend: { fontFamily: {
      sans: ['Poppins', 'sans-serif'],
      serif: ['Playfair Display', 'serif'],
    },},
  },
  plugins: [],
};
