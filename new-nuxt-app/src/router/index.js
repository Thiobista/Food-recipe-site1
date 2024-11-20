import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '@/views/Home.vue'; // Example

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/create-account',
    name: 'CreateAccount',
    component: () => import('@/views/CreateAccount.vue'), // Lazy load
  },
];

const router = new VueRouter({
  mode: 'history', // Use history mode to avoid hash URLs
  routes,
});

export default router;
