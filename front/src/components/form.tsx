import React, { useState } from 'react';
import { postTodoJSON, postTodoText } from '../apis/post';
import { Todo } from '../apis/types';

export const TodoForm = (): JSX.Element => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  // var description: string;
  const [duedate, setDuedate] = useState(new Date());
  return (
    <form
      onSubmit={async (e) => {
        e.preventDefault();
        const todo: Todo = {
          title,
          description,
          duedate,
        };
        try {
          const newTodo = await postTodoJSON(todo);
          console.log(newTodo);
        } catch (e) {
          console.error(e);
        }
      }}
      style={{
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <div>
        <input
          placeholder='title'
          type='text'
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        {title}
      </div>

      <div>
        <input
          placeholder='description'
          type='text'
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          // onChange={(e) => {
          //   description = e.target.value;
          // }}
          // reactive とは何かを説明するためのコメントアウト
        />
        {description}
      </div>

      <div>
        <input
          placeholder='duedate'
          type='datetime-local'
          value={duedate.toISOString().slice(0, 16)}
          onChange={(e) => setDuedate(new Date(e.target.value))}
        />
        {duedate.toString()}
      </div>

      <input type='submit' />
    </form>
  );
};
