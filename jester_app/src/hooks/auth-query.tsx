import type{ DefaultError, QueryKey, UseQueryOptions, UseQueryResult } from "@tanstack/react-query";
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "../components/context/hook";
import { useMessageApi } from "../components/context/message";
import type { MessageInstance } from "antd/es/message/interface";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { AxiosError } from "axios";

const useAuthenticatedQuery = <TQueryFnData = unknown, TError = DefaultError, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey>
    (options: UseQueryOptions<TQueryFnData, TError, TData, TQueryKey> ) : UseQueryResult<TData, TError> => {

    const { refetchAuth } = useAuth();

    const messageApi = useMessageApi() as MessageInstance;
    const { error: errorNotify } = messageApi;

    //router
    const router = useNavigate();

    const results = useQuery({
        retry: false,
        refetchOnMount: true,
        ...options,
    });

    const { error } = results;

    useEffect(() => {
        if (error instanceof AxiosError ){
            if (error.response?.status === 404){
                router("/404");
            }

            if (error.response?.status === 401){
                errorNotify({
                    content: "Session expired. Please log in again.",
                    key: "auth-error",
                });
            
                refetchAuth();
            }

            if (error.response?.status === 500) {
                errorNotify({
                    content: "Server error. Please try again later.",
                    key: "server-error",
                });
            }
        }
    },[error, errorNotify, refetchAuth, router]);

    return results;
}

export default useAuthenticatedQuery;