import { createBrowserRouter } from "react-router-dom";
import Login from "./modules/login";
import AuthenticationLayout from "./modules/authentication";
import DashboardLayout from "./modules";
import Home from "./modules/home";
import Periods from "./modules/periods";

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
                path: "/budgets", //TODO: still needs to be implemented
                element: <div>Budgets Page</div>
            },
            {
                path: "/periods",
                element: <Periods />
            },
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