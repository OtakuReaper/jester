import type React from "react";
import { AuthContext } from "./auth-context";
import { useEffect, useState } from "react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import type { AuthenticationState } from "../../models/auth";
import { getProfile, login, logout } from "../../services/authentication";

export function AuthProvider({   children }: {children: React.ReactNode}) {
    
    //states
    const [auth, setAuth] = useState(null);
    const [loading, setLoading] = useState(true);
    const queryClient = useQueryClient();

    //queries
    const { data, refetch: refetchAuth, status } = useQuery({
        queryKey: ['authentication'],
        queryFn: getProfile,
        refetchInterval: auth ? (1000 * 60 * 5) : false, 
        retry: false,
    })

    //mutation handler functions
    const logoutSuccessHandler = () => {
        refetchAuth();
    }

    const logoutErrorHandler = () => {
        refetchAuth();
    }

    const loginSuccessHandler = () => {
        queryClient.clear();
        refetchAuth();
    }

    //mutations
    const { mutate: logoutMutation, isPending: logoutLoading, error: logoutError } = useMutation({
        mutationFn: logout,
        onError: logoutErrorHandler,
        onSuccess: logoutSuccessHandler,
    })

    const { mutate: loginMutation, isPending: loginLoading, error: loginError } = useMutation({
        mutationFn: login,
        onSuccess: loginSuccessHandler,
    })

    //effects
    useEffect(() => {

        if(data && status == 'success'){
            setAuth(data);
            setLoading(false);
        }

        if(status == 'error'){
            setAuth(null);
            setLoading(false);
        }

    },[data, status, setAuth, setLoading])
    
    const currentState: AuthenticationState = {
        auth,
        loading: loading,
        logoutHandler: logoutMutation,
        logoutLoading,
        logoutError,
        loginHandler: loginMutation,
        loginLoading,
        loginError,
        refetchAuth,
    }
    
    return <AuthContext.Provider value={currentState}>{children}</AuthContext.Provider>
}