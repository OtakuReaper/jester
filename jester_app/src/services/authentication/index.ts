// import axios from "../../utils/singletons/axios";
import rawAxios from 'axios';

const unsafeInstance = rawAxios.create({
    baseURL: import.meta.env.VITE_PUBLIC_API_ORIGIN,
})

export const login = async (data: {username: string; password: string}) => {
    const response = await unsafeInstance.post("/login", data);
    return response.data;
}