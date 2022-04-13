import { gql, useLazyQuery, useQuery } from '@apollo/client';
import { useEffect, useState } from 'react';
import { Todo } from '../apis/models/todo';

const FETCH_ALL_TODOS = gql`
  query fetchAllTodos {
    todos {
      id
      title
      description
      dueDate
      createdAt
    }
  }
`;

const FETCH_TODO_BY_ID = gql`
  query fetchTodoById($id: ID) {
    todo(id: $id) {
      id
      title
      description
      dueDate
      createdAt
    }
  }
`;

const TodoViewer = (): JSX.Element => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [lastUpdatedAt, setLastUpdatedAt] = useState(new Date());
  const [fetchAllTodos, data] = useLazyQuery<Todo[]>(FETCH_ALL_TODOS);
  useEffect(() => {
    const fetchAndSetTodos = () => {
      const newTodos = fetchAllTodos();
      //setTodos([...newTodos]);
      setLastUpdatedAt(new Date());
    };
    fetchAndSetTodos();
    const interval = setInterval(fetchAndSetTodos, 5_000);
    return () => clearInterval(interval);
  }, []);

  return (
    <>
      <div>last-updated: {lastUpdatedAt.toString()}</div>
      <table
        style={{
          border: '1px solid black',
        }}
      >
        <thead>
          <tr>
            <th>Title</th>
            <th>Description</th>
            <th>Due Date</th>
            <th>Created At</th>
          </tr>
        </thead>
        <tbody>
          {todos.map((newTodo, i) => (
            <tr key={newTodo.id ?? i}>
              <td>{newTodo.title}</td>
              <td>{newTodo.description ?? ''}</td>
              <td>{newTodo.dueDate.toString()}</td>
              <td>{newTodo.createdAt.toString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
};

export default TodoViewer;
