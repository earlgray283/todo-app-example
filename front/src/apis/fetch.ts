import { http } from './axios';
import { Todo } from './types';

// application/json で受け取ってそのまま返す
export const fetchTodoJSON = async (id: number): Promise<Todo> => {
  const resp = await http.get<Todo>(`/json/todos/${id}`);
  return resp.data;
};

export const fetchTodoText = async (id: number): Promise<Todo> => {
  const resp = await http.get(`/text/todos/${id}`);
  return JSON.parse(JSON.stringify(resp.data));
};

export const fetchAllTodosJSON = async (): Promise<Todo[]> => {
  const resp = await http.get<Todo[]>('/json/todos');
  return resp.data;
};

export const fetchAllTodosText = async (): Promise<Todo[]> => {
  const resp = await http.get('/text/todos');
  return JSON.parse(JSON.stringify(resp.data));
};
