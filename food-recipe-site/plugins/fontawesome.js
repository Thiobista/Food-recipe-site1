// plugins/fontawesome.js
import { library, config } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

// Import Brand Icons
import { faFacebook, faInstagram, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';

// Import Solid Icons
import { faStar, faStarHalfAlt } from '@fortawesome/free-solid-svg-icons';

import '@fortawesome/fontawesome-svg-core/styles.css'; // Import Font Awesome styles

config.autoAddCss = false; // Disable automatic CSS injection (Nuxt handles it)

// Add desired icons to the library
library.add(faFacebook, faInstagram, faTwitter, faYoutube, faStar, faStarHalfAlt);

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('font-awesome-icon', FontAwesomeIcon);
});
