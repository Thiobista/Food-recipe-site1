<template>
  <div>
    <!-- Recipe Card -->
    <div class="bg-white p-6 shadow-lg rounded-lg hover:shadow-xl transition-shadow duration-300">
      <img :src="image" alt="Recipe Image" class="w-full h-56 object-cover rounded-lg shadow-md" />
      <h3 class="font-bold text-2xl mt-4 text-gray-800">{{ title }}</h3>
      <p class="text-gray-600 text-sm mt-2">{{ description }}</p>

      <div class="flex justify-between mt-6 items-center">
        <!-- Rate Button -->
        <button
          @click="openRatingPopup"
          class="text-yellow-500 hover:text-yellow-600 transition-colors duration-300 text-lg"
        >
          ⭐ Rate
        </button>

        <!-- Bookmark Button -->
        <button
          @click="bookmarkRecipe"
          class="text-blue-500 hover:text-blue-600 transition-colors duration-300 text-lg"
        >
          🔖 Bookmark
        </button>

        <!-- Comment Button -->
        <button
          @click="showCommentFormModal"
          class="text-green-500 hover:text-green-600 transition-colors duration-300 text-lg"
        >
          💬 Comment
        </button>

        <!-- Buy Button -->
        <button
          @click="buyRecipe"
          class="text-orange-500 hover:text-orange-600 transition-colors duration-300 text-lg"
        >
          🛒 Buy
        </button>
      </div>
    </div>

    <!-- Rating Popup (appears when rating is clicked) -->
    <div v-if="showRatingPopup" class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center z-30">
      <div class="bg-white p-6 rounded-lg shadow-xl w-80">
        <h3 class="text-xl font-semibold mb-4 text-gray-800">Select Your Rating</h3>
        <div class="flex justify-center space-x-1">
          <i
            v-for="star in 5"
            :key="star"
            @click="setRating(star)"
            :class="[
              star <= (hoverRating || currentRating) ? 'fas fa-star text-yellow-400' : 'far fa-star text-gray-300',
              'cursor-pointer text-3xl transition-transform transform hover:scale-110'
            ]"
            @mouseover="hoverRating = star"
            @mouseleave="hoverRating = currentRating"
          ></i>
        </div>
        <div class="flex justify-center mt-6 space-x-4">
          <button
            @click="cancelRating"
            class="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400 transition-all"
          >
            Cancel
          </button>
          <button
            @click="confirmRating"
            class="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600 transition-all"
          >
            OK
          </button>
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
          <button
            @click="submitComment"
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-all"
          >
            Submit
          </button>
          <button
            @click="closeModal"
            class="bg-gray-500 text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition-all"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>

    <!-- Modal for Payment with Chapa -->
    <div v-if="isBuying" class="fixed inset-0 bg-gray-800 bg-opacity-70 flex justify-center items-center z-30">
      <div class="bg-white p-6 rounded-lg shadow-xl w-80">
        <h3 class="text-xl font-semibold mb-4 text-gray-800">Continue Payment with Chapa</h3>
        <p class="text-gray-700 mb-4">You are about to proceed with the payment.</p>
        <div class="flex justify-between">
          <button
            @click="proceedToPayment"
            class="bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600 transition-all"
          >
            Continue
          </button>
          <button
            @click="closePaymentModal"
            class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-all"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["image", "title", "description", "recipeId"],
  data() {
    return {
      isBuying: false,
      showComments: false,
      showModal: false,
      commentText: "",
      comments: [],
      currentRating: 0, // Initially no rating
      hoverRating: 0,   // Initially no hover effect
      showRatingPopup: false, // Controls the popup for selecting a rating
      page: 1,
      loadingMore: false,
    };
  },
  created() {
    this.loadComments();
  },
  methods: {
    openRatingPopup() {
      this.showRatingPopup = true;
    },
    closeRatingPopup() {
      this.showRatingPopup = false;
    },
    setRating(star) {
      this.currentRating = star; // Update the rating when a star is clicked
    },
    cancelRating() {
      this.showRatingPopup = false; // Cancel rating and do nothing
    },
    confirmRating() {
      this.showRatingPopup = false; // Confirm the rating and close the popup
    },
    bookmarkRecipe() {
      this.$toast.success("Recipe added to bookmarks!");
    },
    showCommentFormModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    loadComments() {
      const newComments = [
        { text: "Great recipe!" },
        { text: "I loved this! Will make again." },
        { text: "Delicious and easy to follow." },
      ];
      this.comments = [...this.comments, ...newComments];
    },
    handleScroll(event) {
      const bottom = event.target.scrollHeight === event.target.scrollTop + event.target.clientHeight;
      if (bottom && !this.loadingMore) {
        this.loadMoreComments();
      }
    },
    loadMoreComments() {
      this.loadingMore = true;
      setTimeout(() => {
        const additionalComments = [
          { text: "Amazing! Highly recommend." },
          { text: "My family loved it!" },
          { text: "One of the best recipes ever." },
        ];
        this.comments = [...this.comments, ...additionalComments];
        this.loadingMore = false;
      }, 1500);
    },
    submitComment() {
      if (this.commentText.trim()) {
        this.comments.push({ text: this.commentText });
        this.commentText = "";
        this.showModal = false;
      } else {
        this.$toast.error("Please write a comment.");
      }
    },
    buyRecipe() {
      this.isBuying = true;
    },
    proceedToPayment() {
      this.$toast.success("Proceeding to payment...");
      this.isBuying = false;
    },
    closePaymentModal() {
      this.isBuying = false;
    },
  },
};
</script>

<style scoped>
/* Add custom styles here */
</style>
