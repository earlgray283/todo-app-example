import { ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache({ addTypename: false }), // デフォルトだと response に "__typename" field が付加されてしまうので false にする
});

export default client;
