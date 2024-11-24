import { provideApolloClient } from "@vue/apollo-composable";
import { apolloClient } from "~/apollo/client"; // Adjust for named or default export

export default defineNuxtPlugin((nuxtApp) => {
  provideApolloClient(apolloClient);
});
