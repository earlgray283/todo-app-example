import React, { useState } from 'react';
import {
  fetchAllTodosJSON,
  fetchTodoJSON,
  fetchTodoText,
  isAxiosError,
} from './apis/fetch';
import { Todo, TodoRespError } from './apis/types';
import { InputIdButton } from './components/button';
import './styles/App.css';

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  return (
    <div className='App'>
      <div>
        <InputIdButton
          label='Fetch todo(application/json)'
          onClick={async (id: number) => {
            try {
              const todo = await fetchTodoJSON(id);
              setTodos([...todos, todo]);
            } catch (e: unknown) {
              if (isAxiosError<TodoRespError>(e) && e.response) {
                console.error(e.message);
              } else {
                console.error(e);
              }
            }
          }}
        />
        <InputIdButton
          label='Fetch todo(plain/text)'
          onClick={async (id: number) => {
            try {
              const todo = await fetchTodoText(id);
              setTodos([...todos, todo]);
            } catch (e: unknown) {
              if (isAxiosError<TodoRespError>(e) && e.response) {
                console.error(e.message);
              } else {
                console.error(e);
              }
            }
          }}
        />

        <button
          onClick={async () => {
            try {
              const todos = await fetchAllTodosJSON();
              setTodos(todos);
            } catch (e: unknown) {
              if (isAxiosError<TodoRespError>(e) && e.response) {
                console.error(e.message);
              } else {
                console.error(e);
              }
            }
          }}
        >
          Fetch all todos(plain/text)
        </button>
      </div>

      <div>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {todos.map((todo, i) => (
              <tr key={i}>
                <th>{todo.id}</th>
                <th>{todo.detail.title}</th>
                <th>{todo.detail.description}</th>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default App;
