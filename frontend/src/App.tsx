import { getAuth, signOut, User } from 'firebase/auth';
import { createContext, useEffect, useState } from 'react';
import SigninPage from './pages/Signin';
import './lib/firebase';
import TodoHome from './pages/TodoHome';
import { useMutation } from '@apollo/client';
import { SESSION_LOGIN } from './apis/auth';
import { SessionToken } from './apis/models/auth';

interface UserContextType {
  user: User | null;
  setUser: (newUser?: User) => void;
}

const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
});

function App() {
  const [user, setUser] = useState<User | null | undefined>();
  const auth = getAuth();
  const [sessionLogin] = useMutation<{
    sessionLogin: SessionToken;
    token: string;
  }>(SESSION_LOGIN);

  useEffect(() => {
    auth.onAuthStateChanged(async (user) => {
      try {
        if (user) {
          const idToken = await user.getIdToken();
          await sessionLogin({ variables: { token: idToken } });
        }
        setUser(user);
      } catch (e) {
        console.log(e);
      }
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
          <div>
            <button
              onClick={async () => {
                await signOut(auth);
                alert('signout');
              }}
            >
              Sign out
            </button>
            <TodoHome />
          </div>
        ) : user === null ? (
          <SigninPage />
        ) : (
          <div>loading...</div>
        )}
      </div>
    </UserContext.Provider>
  );
}

export default App;
