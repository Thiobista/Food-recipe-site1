<template>
  <div class="recipe-card-container">
    <!-- Recipe Card -->
    <div class="recipe-card">
      <div
  class="shadow-lg rounded-lg hover:shadow-xl hover:scale-105 transition-transform duration-300 flex flex-col bg-red"
>

        <!-- Recipe Image -->
        <div class="relative">
          <img :src="image" alt="Recipe Image" class="w-full h-56 object-cover rounded-lg shadow-md" />
        </div>

        <!-- Recipe Title and Description -->
        <div class="mt-4">
          <h3 class="font-bold text-2xl text-white-800">{{ title }}</h3>
          <p class="text-white-600 text-sm mt-2">{{ description }}</p>
        </div>

      <!-- Action Buttons -->
<div class="p-7 flex justify-between mt-6 items-center">
  <!-- Rate Button -->
  <button
    @click="openRatingPopup"
    class="text-yellow-500 hover:text-yellow-600 transition-colors duration-300 text-lg flex flex-col items-center"
  >
    <span class="text-2xl">⭐</span>
    <span class="text-sm">Rate</span>
  </button>

  <!-- Bookmark Button -->
  <button
    @click="bookmarkRecipe"
    class="text-yellow-500 hover:text-yellow-600 transition-colors duration-300 text-lg flex flex-col items-center"
  >
    <span class="text-2xl">{{ isBookmarked ? '✅' : '🔖' }}</span>
    <span class="text-sm">{{ isBookmarked ? 'Bookmarked' : 'Bookmark' }}</span>
  </button>

  <!-- Comment Button -->
  <button
    @click="showCommentFormModal"
    class="text-yellow-500 hover:text-yellow-600 transition-colors duration-300 text-lg flex flex-col items-center"
  >
    <span class="text-2xl">💬</span>
    <span class="text-sm">Comment</span>
  </button>

  <!-- Buy Button -->
  <button
    @click="buyRecipe"
    class="text-yellow-500 hover:text-yellow-600 transition-colors duration-300 text-lg flex flex-col items-center"
  >
    <span class="text-2xl">🛒</span>
    <span class="text-sm">Buy</span>
  </button>
</div>

      </div>
    </div>

    <!-- Rating Popup -->
    <div v-if="showRatingPopup" class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center z-30">
      <div class="bg-white p-6 rounded-lg shadow-xl w-80">
        <h3 class="text-xl font-semibold mb-4 text-gray-800">Select Your Rating</h3>
        <div class="flex justify-center space-x-1">
          <i
  v-for="star in 5"
  :key="star"
  @click="setRating(star)"
  :class="[
    star <= (hoverRating || currentRating) ? 'fas fa-star text-yellow-400' : 'far fa-star text-gray-600', // Update this class for better contrast
    'cursor-pointer text-3xl transition-transform transform hover:scale-110'
  ]"
  @mouseover="hoverRating = star"
  @mouseleave="hoverRating = currentRating"
/>

        </div>
        <div class="flex justify-center mt-6 space-x-4">
          <button @click="cancelRating" class="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400 transition-all">Cancel</button>
          <button @click="confirmRating" class="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600 transition-all">OK</button>
        </div>
      </div>
    </div>

    <!-- Comments Section -->
    <div v-if="showComments" class="mt-8">
      <h3 class="text-2xl font-semibold mb-4 text-gray-800">Comments</h3>
      <div class="comments-container max-h-[300px] overflow-y-auto p-2 bg-gray-50 rounded-lg">
        <div v-for="(comment, index) in comments" :key="index" class="mb-4 p-4 bg-gray-100 rounded-lg shadow-sm hover:shadow-md transition-shadow duration-200">
          <p class="text-gray-700 text-sm">{{ comment.text }}</p>
        </div>
      </div>
      <div v-if="loadingMore" class="text-center mt-6">
        <button class="bg-blue-500 text-white px-4 py-2 rounded-lg">Loading...</button>
      </div>
    </div>

    <!-- Modal for Comment Form -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center z-30">
      <div class="bg-white p-6 rounded-lg shadow-xl w-80">
        <h3 class="text-xl font-semibold mb-4 text-gray-800">Add Your Comment</h3>
        <div v-if="comments.length > 0" class="mb-4">
          <h4 class="text-lg font-medium mb-2 text-gray-700">Existing Comments:</h4>
          <div v-for="(comment, index) in comments" :key="index" class="p-2 mb-2 bg-gray-200 rounded-lg shadow-sm">
            <p class="text-gray-700 text-sm">{{ comment.text }}</p>
          </div>
        </div>
        <textarea v-model="commentText" class="w-full h-24 p-3 border rounded-lg mb-4 bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="Write your comment here..."></textarea>
        <div class="flex justify-between">
          <button @click="submitComment" class="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600 transition-all">Submit</button>
          <button @click="closeModal" class="bg-gray-500 text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition-all">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Modal for Payment with Chapa -->
    <div v-if="isBuying" class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center z-30">
      <div class="bg-white p-6 rounded-lg shadow-xl w-80">
        <h3 class="text-xl font-semibold mb-4 text-gray-800">Continue Payment with Chapa</h3>
        <p class="text-gray-700 mb-4">You are about to proceed with the payment.</p>
        <div class="flex justify-between">
          <button @click="proceedToPayment" class="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600 transition-all">Continue</button>
          <button @click="closePaymentModal" class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-all">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Notification -->
    <div v-if="notification" class="fixed top-4 right-4 bg-yellow-500 text-white px-4 py-2 rounded-lg shadow-lg z-40 transition-all">
      {{ notification }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const props = defineProps({
  image: String,
  title: String,
  description: String,
  recipeId: String,
});

const isBookmarked = ref(false); // Reactive state for bookmark
const isBuying = ref(false);
const showComments = ref(false);
const showModal = ref(false);
const commentText = ref('');
const comments = ref([]);
const currentRating = ref(0); 
const hoverRating = ref(0);   
const showRatingPopup = ref(false);
const loadingMore = ref(false);
const notification = ref('');

const openRatingPopup = () => showRatingPopup.value = true;
const cancelRating = () => showRatingPopup.value = false;
const confirmRating = () => showRatingPopup.value = false;

const setRating = (star) => {
  currentRating.value = star;
};

const showCommentFormModal = () => {
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const bookmarkRecipe = () => {
  isBookmarked.value = !isBookmarked.value; // Toggle bookmark state
};
const loadComments = () => {
  const newComments = [
    { text: "Great recipe!" },
    { text: "I loved this! Will make again." },
    { text: "Delicious and easy to follow." },
  ];
  comments.value = [...comments.value, ...newComments];
};

const handleScroll = (event) => {
  const bottom = event.target.scrollHeight === event.target.scrollTop + event.target.clientHeight;
  if (bottom && !loadingMore.value) {
    loadMoreComments();
  }
};

const loadMoreComments = () => {
  loadingMore.value = true;
  setTimeout(() => {
    const additionalComments = [
      { text: "Amazing! Highly recommend." },
      { text: "My family loved it!" },
      { text: "One of the best recipes ever." },
    ];
    comments.value = [...comments.value, ...additionalComments];
    loadingMore.value = false;
  }, 1500);
};

const submitComment = () => {
  if (commentText.value.trim()) {
    comments.value.push({ text: commentText.value });
    commentText.value = '';
    closeModal();
  }
};

const buyRecipe = () => {
  isBuying.value = true;
};

const proceedToPayment = () => {
  isBuying.value = false; // Close the payment modal
  notification.value = "Redirecting to Chapa for payment...";
  setTimeout(() => {
    notification.value = '';
    window.location.href = "https://chapa.co"; // Redirect to Chapa website
  }, 3000);
};



const closePaymentModal = () => {
  isBuying.value = false;
};

onMounted(() => {
  loadComments();
});

</script>

<style scoped>
/* Add custom styling as needed */
.text-gray-600 {
  color: #6b7280; /* This is darker gray */
}


</style>   