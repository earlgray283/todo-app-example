import { getAuth, User } from 'firebase/auth';
import { createContext, useEffect, useState } from 'react';
import SigninPage from './pages/Signin';
import './apis/firebase';
import TodoHome from './pages/TodoHome';

interface UserContextType {
  user: User | null;
  setUser: (newUser?: User) => void;
}

const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
});

function App() {
  const [user, setUser] = useState<User | null>();
  const auth = getAuth();

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      setUser(user);
    });
  }, []);

  return (
    <UserContext.Provider
      value={{
        user: auth.currentUser,
        setUser: (newUser) => setUser(newUser),
      }}
    >
      <div className='App'>
        {user ? (
          <TodoHome />
        ) : user == null ? (
          <SigninPage />
        ) : (
          <div>loading...</div>
        )}
      </div>
    </UserContext.Provider>
  );
}

export default App;
