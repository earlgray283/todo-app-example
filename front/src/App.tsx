import React from 'react';
import {
  fetchAllTodosJSON,
  fetchAllTodosText,
  fetchTodoJSON,
  fetchTodoText,
} from './apis/fetch';
import { FetchAllTodosButton, FetchTodoButton } from './components/button';
import { TodoForm } from './components/form';
import './styles/App.css';

function App() {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
      }}
    >
      <TodoForm />

      <FetchTodoButton
        label='Fetch todo(application/json)'
        onClick={fetchTodoJSON}
      />
      <FetchTodoButton label='Fetch todo(plain/text)' onClick={fetchTodoText} />

      <FetchAllTodosButton
        label='Fetch all todos(application/json)'
        onClick={fetchAllTodosJSON}
      />
      <FetchAllTodosButton
        label='Fetch all todos(plain/text)'
        onClick={fetchAllTodosText}
      />
    </div>
  );
}

export default App;
