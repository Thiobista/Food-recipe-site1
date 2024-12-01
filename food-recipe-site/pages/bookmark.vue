<template>
  <div class="fufu">
    <Header />

    <!-- Bookmarks Section -->
    <main class="container mx-auto py-16 mt-12">
      <h1 class="text-3xl font-bold text-center mb-8">Your Bookmarked Recipes</h1>
      <div v-if="bookmarkedRecipes.length === 0" class="text-center text-gray-600">
        <p>You don't have any bookmarked recipes yet.</p>
        <button
          class="mt-4 bg-yellow-500 text-white py-2 px-6 rounded-lg shadow-lg hover:bg-yellow-600 transition duration-300"
          @click="goToPage('/recipes')"
        >
          <i class="fas fa-search"></i> Explore Recipes
        </button>
      </div>

      <div v-else class="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div
  v-for="(recipe, index) in bookmarkedRecipes"
  :key="index"
  class="recipe-card bg-white rounded-lg shadow-lg overflow-hidden relative flex flex-col"
>
  <img :src="recipe.image" :alt="recipe.title" class="w-full h-40 object-cover" />
  <div class="p-4 flex-grow">
    <h2 class="text-xl font-bold">{{ recipe.title }}</h2>
    <p class="text-gray-600 text-sm mb-4">{{ recipe.description }}</p>
  </div>
  <div class="action-bar bg-gray-100 p-2 flex justify-around border-t border-gray-200">
    <button
      class="flex flex-col items-center text-red-500 hover:text-red-600"
      @click="removeBookmark(recipe.id)"
      aria-label="Remove Bookmark"
    >
      <font-awesome-icon icon="trash" class="text-xl" />
      <span class="text-sm">Remove</span>
    </button>
    <button
      class="flex flex-col items-center text-orange-500 hover:text-orange-600"
      @click="shareRecipe(recipe)"
      aria-label="Share Recipe"
    >
      <font-awesome-icon icon="share" class="text-xl" />
      <span class="text-sm">Share</span>
    </button>
    <button
      class="flex flex-col items-center text-orange-500 hover:text-orange-600"
      @click="openRateModal(recipe.id)"
      aria-label="Rate Recipe"
    >
      <font-awesome-icon icon="star" class="text-xl" />
      <span class="text-sm">Rate</span>
    </button>
    <button
      class="flex flex-col items-center text-orange-500 hover:text-orange-600"
      @click="openCommentModal(recipe.id)"
      aria-label="Comment on Recipe"
    >
      <font-awesome-icon icon="comment" class="text-xl" />
      <span class="text-sm">Comment</span>
    </button>
  </div>
</div>

      </div>

      <!-- Rate Modal -->
<transition name="fade">
  <div v-if="showRateModal" class="modal">
    <div class="modal-content max-w-md p-6 bg-white rounded-lg shadow-lg relative">
      <h2 class="text-2xl font-bold mb-4">Rate Recipe</h2>
      <p class="mb-4">Rate <span class="font-semibold">{{ selectedRecipe?.title }}</span></p>
      <div class="flex justify-center mb-4">
        <span
          v-for="star in 5"
          :key="star"
          class="cursor-pointer text-4xl"
          :class="{
            'text-yellow-500': star <= starRating,
            'text-gray-400': star > starRating,
          }"
          @click="setRating(star)"
        >
          ★
        </span>
      </div>
      <div class="flex justify-end gap-4">
        <button
          class="py-2 px-4 bg-yellow-500 text-white rounded-lg hover:bg-yellow-600"
          @click="submitRating"
        >
          Submit
        </button>
        <button
          class="py-2 px-4 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400"
          @click="closeModals"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</transition>


     
    <!-- Comment Modal -->
<transition name="fade">
  <div v-if="showCommentModal" class="modal">
    <div class="modal-content max-w-md p-6 bg-white rounded-lg shadow-lg relative">
      <h2 class="text-2xl font-bold mb-4">Comments for {{ selectedRecipe?.title }}</h2>

      <!-- Existing Comments Section -->
      <div class="comments-section mb-4">
        <h3 class="text-lg font-semibold mb-2">Existing Comments:</h3>
        <div v-if="selectedRecipe?.comments?.length">
          <ul class="space-y-2">
            <li v-for="(comment, index) in selectedRecipe.comments" :key="index" class="bg-gray-100 p-3 rounded-lg">
              {{ comment }}
            </li>
          </ul>
        </div>
        <p v-else class="text-gray-500">No comments yet. Be the first to comment!</p>
      </div>

      <!-- Add New Comment Section -->
      <textarea
        v-model="comment"
        class="input mb-4 border-gray-300 focus:ring-2 focus:ring-yellow-500"
        placeholder="Write your comment here..."
      ></textarea>
      <div class="flex justify-end gap-4">
        <button
          class="py-2 px-4 bg-yellow-500 text-white rounded-lg hover:bg-yellow-600"
          @click="submitComment"
        >
          Submit
        </button>
        <button
          class="py-2 px-4 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400"
          @click="closeModals"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</transition>

    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const bookmarkedRecipes = ref([
  {
    id: 1,
    title: 'Doro Wot',
    description: 'Spicy chicken stew with hard-boiled eggs.',
    image: '/images/doro.jpg',
    comments: [], // Add this
  },
  {
    id: 2,
    title: 'Shiro',
    description: 'Spiced chickpea stew.',
    image: '/images/shiro.jpg',
    comments: [], // Add this
  },
]);


const router = useRouter();
const goToPage = (route) => {
  router.push(route);
};
const showRateModal = ref(false);
const showCommentModal = ref(false);
const selectedRecipe = ref(null);
const rating = ref('');
const comment = ref('');

const removeBookmark = (id) => {
  bookmarkedRecipes.value = bookmarkedRecipes.value.filter((recipe) => recipe.id !== id);
};

const openRateModal = (id) => {
  selectedRecipe.value = bookmarkedRecipes.value.find((recipe) => recipe.id === id);
  showRateModal.value = true;
};

const submitRating = () => {
  console.log(`Rated ${selectedRecipe.value.title}: ${starRating.value} stars`);
  closeModals();
};

const openCommentModal = (id) => {
  selectedRecipe.value = bookmarkedRecipes.value.find((recipe) => recipe.id === id);
  showCommentModal.value = true;
};

const submitComment = () => {
  if (!comment.value.trim()) return; // Prevent empty comments
  selectedRecipe.value.comments.unshift(comment.value.trim()); // Add new comment to the top
  comment.value = ''; // Clear the input
  console.log(`Commented on ${selectedRecipe.value.title}: ${selectedRecipe.value.comments}`);
};

const closeModals = () => {
  showRateModal.value = false;
  showCommentModal.value = false;
  selectedRecipe.value = null;
  rating.value = '';
  comment.value = '';
};

const shareRecipe = (recipe) => {
  const shareData = {
    title: recipe.title,
    text: recipe.description,
    url: window.location.href,
  };
  if (navigator.share) {
    navigator.share(shareData).catch((err) => console.error(err));
  } else {
    alert('Sharing not supported.');
  }
};
const starRating = ref(0);

const setRating = (rating) => {
  starRating.value = rating;
};



</script>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  width: 90%;
  max-width: 500px;
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  animation: slideDown 0.3s ease-in-out;
}

.text-gray-400 {
  color: #d1d5db; /* Tailwind's gray-400 */
}

.text-yellow-500 {
  color: #f59e0b; /* Tailwind's yellow-500 */
}

.cursor-pointer {
  cursor: pointer;
}

.input {
  width: 100%;
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #ddd;
  outline: none;
}

@keyframes slideDown {
  from {
    transform: translateY(-20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>
