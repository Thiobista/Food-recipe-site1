// plugins/fontawesome.js
import { library, config } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faStar, faStarHalfAlt } from '@fortawesome/free-solid-svg-icons'; // Import necessary icons
// Import Brand Icons
import { faFacebook, faInstagram, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';
import { faEdit, faTrash, faShareAlt } from '@fortawesome/free-solid-svg-icons';
// Import Solid Icons
import { faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons';

import { faStar as fasStar } from '@fortawesome/free-solid-svg-icons'; // Solid star
import { faStar as farStar } from '@fortawesome/free-regular-svg-icons';
import '@fortawesome/fontawesome-svg-core/styles.css'; // Import Font Awesome styles

config.autoAddCss = false; // Disable automatic CSS injection (Nuxt handles it)

// Add desired icons to the library
library.add(faFacebook, faInstagram, faTwitter, faYoutube, faEdit, faTrash, faShareAlt, fasStar, farStar);
library.add(faChevronUp, faChevronDown);
library.add(faStar, faStarHalfAlt);
export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('font-awesome-icon', FontAwesomeIcon);
});
