import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot, useSetRecoilState } from 'recoil'
import { currentUserState } from '../stores/currentUser';
import { useEffect } from 'react';
import { useCurrentUser } from '../hooks/useCurrentUser';
import { fetchCurrentUser } from '../utils/fetchCurrentUser';
import { CurrentUser } from '../types/user';

/* recoil の atom にログイン済みユーザ情報をセット */
const AppInit = (): null => {
  // recoil atom setter
  const setCurrentUser = useSetRecoilState(currentUserState)

  // fetch current user from cookies
  const initUser = async () => {
    const currentUser: CurrentUser | null = await fetchCurrentUser()
    setCurrentUser(currentUser)
  }

  useEffect(() => { initUser() }, [])
  return null
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <Component {...pageProps} />
      <AppInit />
    </RecoilRoot>
  )
}

export default MyApp
