import { gql } from '@apollo/client';

export const CREATE_TODO = gql`
  mutation createTodo($newTodo: NewTodo!) {
    createTodo(input: $newTodo) {
      userId
      title
      description
      dueDate
      createdAt
    }
  }
`;

export const FETCH_ALL_TODOS = gql`
  query fetchAllTodos {
    todos {
      userId
      title
      description
      dueDate
      done
      createdAt
    }
  }
`;
