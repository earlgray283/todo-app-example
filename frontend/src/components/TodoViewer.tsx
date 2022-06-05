import { useLazyQuery } from '@apollo/client';
import { useEffect, useState } from 'react';
import { Todo } from '../apis/models/todo';
import { FETCH_ALL_TODOS } from '../apis/todo';
import './common.css';

const TodoViewer = (): JSX.Element => {
  const [lastUpdatedAt, setLastUpdatedAt] = useState(new Date());

  // useMutation() みたいに [実行関数, 取得した内容] で受け取りたい時は useLazyQuery を使う
  // ここで、useLazyQuery の型引数は必ず interface 型にしないといけない(<Todo[]> はうまくいかない)
  const [fetchAllTodos, { data }] = useLazyQuery<{ todos: Todo[] }>(
    FETCH_ALL_TODOS
  );

  useEffect(() => {
    fetchAllTodos();
    // eslint-disable-next-line react-hooks/exhaustive-deps
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
              <tr key={newTodo.userId ?? i}>
                <td>{newTodo.title}</td>
                <td>{newTodo.description ?? ''}</td>
                <td>{newTodo.dueDate.toString()}</td>
                <td>{newTodo.createdAt.toString()}</td>
              </tr>
            ))}
        </tbody>
      </table>
      <button
        onClick={() => {
          fetchAllTodos();
          setLastUpdatedAt(new Date());
        }}
      >
        Update
      </button>
    </>
  );
};

export default TodoViewer;
