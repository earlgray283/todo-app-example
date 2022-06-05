import { gql } from '@apollo/client';

export const SESSION_LOGIN = gql`
  mutation sessionLogin($token: String!) {
    sessionLogin(token: $token) {
      token
    }
  }
`;
