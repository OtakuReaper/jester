import { createBrowserRouter } from "react-router-dom";
import Login from "./modules/login";
import AuthenticationLayout from "./modules/authentication";
import DashboardLayout from "./modules";
import Home from "./modules/home";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <DashboardLayout />,
        children: [
            {
                path: "/",
                element: <Home/>
            },
            {
                path: "/budgets",
                element: <div>Budgets Page</div>
            }
        ]
    },
    {
        path: "/auth",
        element: <AuthenticationLayout />,
        children: [
            {
                path: "login",
                element: <Login />
            }
        ]
    }
]);