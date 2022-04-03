export interface Todo {
  id: number;
  detail: TodoDetail;
}

export interface TodoDetail {
  title: string;
  description: string;
}

export interface TodoRespError {
  status: number;
  message: string;
}
