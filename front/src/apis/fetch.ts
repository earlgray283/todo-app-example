import axios, { AxiosError } from 'axios';
import { Todo } from './types';

export const http = axios.create({ baseURL: 'http://localhost:8081' });

export function isAxiosError<T = any, D = any>(e: any): e is AxiosError<T, D> {
  return !!e.isAxiosError;
}

// application/json で受け取ってそのまま返す
export const fetchTodoJSON = async (id: number): Promise<Todo> => {
  const resp = await http.get<Todo>('/todos', { params: { id: id } });
  return resp.data;
};

// plain/text で受け取って json に parse して返す
export const fetchTodoText = async (id: number): Promise<Todo> => {
  const resp = await http.get('/todos', { params: { id: id } });
  // const data: string = resp.data;  // NG: resp.data は Object 型だから json 文字列に変換はできない
  const data = JSON.stringify(resp.data);
  return JSON.parse(data);
};

export const fetchAllTodosText = async (): Promise<Todo[]> => {
  const resp = await http.get('/todos');
  const data = JSON.stringify(resp.data);
  return JSON.parse(data);
};

export const fetchAllTodosJSON = async (): Promise<Todo[]> => {
  const resp = await http.get<Todo[]>('/todos');
  return resp.data;
};
