import { Todo } from './types';
import { http } from './axios';

export async function postTodoJSON(todo: Todo): Promise<Todo> {
  const resp = await http.post<Todo>('/json/todos', todo);
  return resp.data;
}

export async function postTodoText(todo: Todo): Promise<Todo> {
  const jsonText = JSON.stringify(todo);
  const resp = await http.post<Todo>('/text/todos', jsonText);
  return resp.data;
}
