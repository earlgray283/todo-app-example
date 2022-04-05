import axios, { AxiosError } from 'axios';

export const http = axios.create({
  baseURL: 'http://localhost:8081',
});


export function isAxiosError<T = any, D = any>(e: any): e is AxiosError<T, D> {
  return !!e.isAxiosError;
}
