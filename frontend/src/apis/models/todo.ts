export interface Todo {
  userId?: string;
  title: string;
  description?: string;
  dueDate: string;
  done?: boolean;
  createdAt: string;
  updatedAt?: string;
}

export interface NewTodo {
  title: string;
  description?: string;
  dueDate: string;
}
