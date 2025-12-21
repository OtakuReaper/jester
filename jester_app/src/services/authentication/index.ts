// import axios from "../../utils/singletons/axios";
import rawAxios from 'axios';
import axios from '../../utils/singletons/axios';
import type { Creds } from '../../models/auth';

const unsafeInstance = rawAxios.create({
    baseURL: import.meta.env.VITE_PUBLIC_API_ORIGIN,
})

export const login = async (data: Creds) => {
    const response = await unsafeInstance.post("/auth/login", data);
    return response.data;
}

export const logout = async () => {
    const response = await axios.get("/auth/logout", { withCredentials: true });
    return response.data;
}

export const getProfile = async () => {
    const reponse = await axios.get("/auth/profile", { withCredentials: true });
    return reponse.data;
}