import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client';

const httpLink = createHttpLink({
  uri: `${process.env.REACT_APP_BACKEND_URL}/query`,
  credentials: 'include',
});
const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache({ addTypename: false }), // デフォルトだと response に "__typename" field が付加されてしまうので false にする
});

console.log(process.env.REACT_APP_BACKEND_URL);

export default client;
