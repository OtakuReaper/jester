import type { User } from "./user"

export interface Creds  {
    username: string,
    password: string
}

export interface AuthenticationState {
    auth: User | null,
    loading: boolean,   
    logoutHandler: () => void,
    logoutLoading: boolean,
    logoutError: Error | null,
    loginHandler: (creds: Creds) => void,
    loginLoading: boolean,
    loginError: Error | null,
    refetchAuth: () => void
}