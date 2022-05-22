import { ApolloClient, InMemoryCache } from '@apollo/client';

const client = new ApolloClient({
  uri: `${process.env.REACT_APP_BACKEND_URL}/query`,
  cache: new InMemoryCache({ addTypename: false }), // デフォルトだと response に "__typename" field が付加されてしまうので false にする
});

console.log(process.env.REACT_APP_BACKEND_URL);

export default client;
