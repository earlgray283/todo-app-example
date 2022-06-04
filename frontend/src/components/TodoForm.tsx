import { gql, useMutation } from '@apollo/client';
import React, { useState } from 'react';
import { NewTodo, Todo } from './../apis/models/todo';

const CREATE_TODO = gql`
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

// date -> "0170-07-31T22:00:00"
function toDatetimeLocalFormat(date: Date): string {
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(
    2,
    '0'
  )}-${date.getDate()}T${String(date.getHours()).padStart(2, '0')}:${String(
    date.getMinutes()
  ).padStart(2, '0')}`;
}

const TodoForm = (): JSX.Element => {
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>();
  const [dueDate, setDueDate] = useState<Date>(new Date());

  const [createTodo, { error, data }] = useMutation<
    { createTodo: Todo },
    { newTodo: NewTodo }
  >(CREATE_TODO);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        createTodo({
          variables: { newTodo: { title, description, dueDate } },
        });
      }}
    >
      {error && (
        <div
          style={{
            backgroundColor: '#f8d7da',
            borderColor: '#f5c6cb',
            borderRadius: '0.25rem',
          }}
        >
          error:
          {error?.message ?? 'hoge'}
        </div>
      )}
      {data && (
        <div
          style={{
            backgroundColor: '#d4edda',
            borderColor: '#c3e6cb',
            borderRadius: '0.25rem',
          }}
        >
          request "postTodo" was successful!
          <p>id: {data.createTodo.userId ?? 'userId is null'}</p>
          <p>title: {data.createTodo.title}</p>
          <p>{data.createTodo.description ?? ''}</p>
          <p>{data.createTodo.createdAt.toDateString()}</p>
        </div>
      )}
      <p>
        <label>title: </label>
        <input
          type='text'
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </p>
      <p>
        <label>description(optional): </label>
        <input
          type='text'
          value={description ?? ''}
          onChange={(e) => setDescription(e.target.value)}
        />
      </p>
      <p>
        <label>due date: </label>
        <input
          type='datetime-local'
          value={toDatetimeLocalFormat(dueDate)}
          onChange={(e) => setDueDate(new Date(e.target.value))}
          required
        />
      </p>
      <p>
        <input type='submit' value='submit' />
      </p>
    </form>
  );
};

export default TodoForm;
