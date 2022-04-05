import React, { useState } from 'react';
import { Todo } from '../apis/types';
import { TodoTable } from './table';

export const FetchTodoButton = (props: {
  label: string;
  onClick: (id: number) => Promise<Todo>;
}): JSX.Element => {
  const [id, setID] = useState(0);
  const [todo, setTodo] = useState<Todo | null>(null);
  return (
    <div>
      <input
        name='id'
        type='number'
        value={id}
        onChange={(e) => setID(parseInt(e.target.value, 10))}
      />
      <button
        onClick={async () => {
          try {
            const newTodo = await props.onClick(id);
            setTodo(newTodo);
          } catch (e) {
            console.error(e);
          }
        }}
      >
        {props.label}
      </button>

      <TodoTable todos={todo ? [todo] : []} />
    </div>
  );
};

export const FetchAllTodosButton = (props: {
  label: string;
  onClick: () => Promise<Todo[]>;
}): JSX.Element => {
  const [todos, setTodos] = useState<Todo[]>([]);
  return (
    <div>
      <button
        onClick={async () => {
          try {
            const newTodos = await props.onClick();
            setTodos([...newTodos]);
          } catch (e) {
            console.error(e);
          }
        }}
      >
        {props.label}
      </button>

      <TodoTable todos={todos} />
    </div>
  );
};
