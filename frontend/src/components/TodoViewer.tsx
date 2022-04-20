import { gql, useLazyQuery } from '@apollo/client';
import { useEffect, useState } from 'react';
import { Todo } from '../apis/models/todo';
import './common.css';

const INTERVAL_FETCH_TODO = 30_000; // 30s

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
  const [lastUpdatedAt, setLastUpdatedAt] = useState(new Date());

  // useMutation() みたいに [実行関数, 取得した内容] で受け取りたい時は useLazyQuery を使う
  // ここで、useLazyQuery の型引数は必ず interface 型にしないといけない(<Todo[]> はうまくいかない)
  const [fetchAllTodos, { data }] = useLazyQuery<{ todos: Todo[] }>(
    FETCH_ALL_TODOS
  );

  useEffect(() => {
    fetchAllTodos();
    const interval = setInterval(() => {
      fetchAllTodos();
      setLastUpdatedAt(new Date());
    }, INTERVAL_FETCH_TODO);
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
          {data &&
            data.todos.map((newTodo, i) => (
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
