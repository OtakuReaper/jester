import { createContext } from "react"
import type { AuthenticationState } from "../../models/auth"

const initialState: AuthenticationState = {
    auth: null,
    loading: true,
    logoutHandler: () => {},
    logoutLoading: false,
    logoutError: null,
    loginHandler: () => {},
    loginLoading: false,
    loginError: null,
    refetchAuth: () => {}
}

export const AuthContext = createContext(initialState)