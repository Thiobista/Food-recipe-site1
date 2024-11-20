import Vue from 'vue';
import Toasted from 'vue-toasted';

if (process.client) {
  Vue.use(Toasted, { /* options */ });
}
