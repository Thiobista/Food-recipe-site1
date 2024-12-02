import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core';
import { provideApolloClient } from '@vue/apollo-composable';

export default defineNuxtPlugin(() => {
  const apolloClient = new ApolloClient({
    link: new HttpLink({ uri: 'https://your-hasura-endpoint/v1/graphql', headers: {
      'x-hasura-admin-secret': 'your-admin-secret',
    }}),
    cache: new InMemoryCache(),
  });

  provideApolloClient(apolloClient);
});
