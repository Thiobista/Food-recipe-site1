<template>
  <div class="relative">
    <!-- Header Section -->
    <header class="fixed top-0 left-0 w-full bg-white shadow-md z-50">
      <Header />
    </header>

    <!-- Main Content -->
    <div class="container mx-auto pt-20 py-10">
      <h1 class="text-4xl font-bold text-yellow-500 text-center">Privacy Policy</h1>
      <p class="text-gray-700 mt-4 text-center">
        At RecipeSite, we value your privacy and are committed to protecting your
        personal information. This Privacy Policy outlines how we collect, use,
        and safeguard your data.
      </p>

      <!-- Interactive Accordion -->
      <div class="mt-10 max-w-3xl mx-auto">
        <div v-for="(section, index) in sections" :key="index" class="mb-4">
          <button
            @click="toggleSection(index)"
            class="w-full text-left px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 transition flex justify-between items-center"
          >
            <span class="font-semibold text-gray-700">{{ section.title }}</span>
            <span>
              <font-awesome-icon
                :icon="activeSection === index ? 'chevron-up' : 'chevron-down'"
              />
            </span>
          </button>
          <div
            v-if="activeSection === index"
            class="p-4 bg-gray-100 text-gray-600 rounded-b"
          >
            <p>{{ section.content }}</p>
          </div>
        </div>
      </div>

      <!-- Accept Button -->
      <div class="text-center mt-8">
        <button
          @click="acceptPolicy"
          class="px-6 py-2 bg-yellow-500 text-white rounded hover:bg-yellow-600 transition focus:outline-none"
        >
          Accept
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { useToast } from "vue-toastification";
import { useRouter } from "vue-router";

export default {
  name: "PrivacyPolicy",
  data() {
    return {
      activeSection: null,
      sections: [
        {
          title: "What Information We Collect",
          content:
            "We collect information you provide directly to us, such as your name, email address, and any other details submitted through forms on our site.",
        },
        {
          title: "How We Use Your Information",
          content:
            "Your information is used to provide and improve our services, process your inquiries, and send updates about RecipeSite.",
        },
        {
          title: "How We Protect Your Data",
          content:
            "We implement a variety of security measures to ensure the safety of your personal information. Access to your data is restricted to authorized personnel only.",
        },
        {
          title: "Your Rights",
          content:
            "You have the right to access, update, or delete your personal information. Contact us at privacy@recipesite.com for any concerns.",
        },
        {
          title: "Policy Updates",
          content:
            "We may update this Privacy Policy from time to time. Please review this page periodically for any changes.",
        },
      ],
    };
  },
  setup() {
    const toast = useToast();
    const router = useRouter();

    const acceptPolicy = () => {
      toast.success("Thank you for accepting the Privacy Policy!", {
        timeout: 3000,
        position: "top-right",
      });

      // Redirect to home page after a delay
      setTimeout(() => {
        router.push("/");
      }, 3000); // 3-second delay
    };

    return { acceptPolicy };
  },
  methods: {
    toggleSection(index) {
      this.activeSection = this.activeSection === index ? null : index;
    },
  },
};
</script>

<style scoped>
button {
  outline: none;
}
</style>
