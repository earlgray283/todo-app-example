import React from 'react';
import { Todo } from '../apis/types';

export const TodoTable = (props: { todos: Todo[] }): JSX.Element => (
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Title</th>
        <th>Description</th>
        <th>Duedate</th>
      </tr>
    </thead>
    <tbody>
      {props.todos.map((todo, i) => (
        <tr key={i}>
          <th>{todo.id}</th>
          <th>{todo.title}</th>
          <th>{todo.description}</th>
          <th>{todo.duedate}</th>
        </tr>
      ))}
    </tbody>
  </table>
);
