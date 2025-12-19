import { createBrowserRouter } from "react-router-dom";
import Login from "./modules/login";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Login />
    }
]);