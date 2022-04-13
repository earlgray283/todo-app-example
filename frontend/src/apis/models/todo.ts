export interface Todo {
  id?: string;
  title: string;
  description?: string;
  dueDate: Date;
  done?: boolean;
  createdAt: Date;
  updatedAt?: Date;
}

export interface NewTodo {
  title: string;
  description?: string;
  dueDate: Date;
}
