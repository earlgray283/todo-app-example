import TodoForm from '../components/TodoForm';
import TodoViewer from '../components/TodoViewer';

function TodoHome(): JSX.Element {
  return (
    <div>
      <TodoForm />
      <TodoViewer />
    </div>
  );
}

export default TodoHome;
