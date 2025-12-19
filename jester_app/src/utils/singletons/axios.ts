import axios from 'axios';

const origin = import.meta.env.VITE_PUBLIC_API_ORIGIN

const instance = axios.create({
    baseURL: origin,
    withCredentials: true,
})

export default instance;