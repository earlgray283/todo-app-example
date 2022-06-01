import React from 'react';
import {
  browserSessionPersistence,
  getAuth,
  getRedirectResult,
  GithubAuthProvider,
  signInWithRedirect,
} from 'firebase/auth';

function SigninWithGithub(): JSX.Element {
  const provider = new GithubAuthProvider();
  const auth = getAuth();
  auth.setPersistence(browserSessionPersistence);
  const signin = async () => {
    try {
      await signInWithRedirect(auth, provider);
      const res = await getRedirectResult(auth);
      if (res) {
        console.log(res.user);
      }
    } catch (e: unknown) {
      console.error(e);
    }
  };

  return <button onClick={signin}>Sign in with Github</button>;
}

export default SigninWithGithub;
