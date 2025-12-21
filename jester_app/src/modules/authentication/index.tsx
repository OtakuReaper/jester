import { Outlet, useNavigate } from "react-router-dom"
import { useAuth } from "../../components/context/hook"
import { useEffect } from "react";

const AuthenticationLayout = () => {

    const { auth, loading } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        if (!loading && auth){
            navigate("/");
        }
    }, [auth, loading, navigate]);

    if(loading){
        return <div> Loading... </div> //TODO: implement a better loading screen
    }

    return (
        <Outlet />
    )
}

export default AuthenticationLayout;