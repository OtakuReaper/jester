import { Layout, Menu, Typography } from "antd";
import { Content, Footer, Header, } from "antd/es/layout/layout";
import { Outlet, useNavigate } from "react-router-dom"
import { useAuth } from "../components/context/hook";
import { useEffect } from "react";


const DashboardLayout = () => {


    //adding authentication s check
    const { auth, loading, } = useAuth(); //TODO: implement logout functionality
    const navigate = useNavigate();
 
    useEffect(() => {
        if (!auth && !loading) {
            navigate("/auth/login");
        }
    },[auth, loading, navigate])


    //menu items
    const menuItems = [
        {
            key: "1",
            label: "Home",
            href: "/"
        },
        {
            key: "2",
            label: "Budgets",
            href: "/budgets"
        }
    ]

    if (loading || !auth) {
        return <div>Loading...</div>; //TODO: Replace with a proper loading spinner
    }

    return (
        <Layout style={{ minHeight: "100vh" }}>
            <Header style={{
                display: "flex",
                alignItems: "center",
                backgroundColor: "var(--secondary-color)",
            }}>
                <Typography.Title style={{ 
                    color: "white", 
                    margin: 0, 
                    textAlign: "center" 
                    }} level={3}>Jester</Typography.Title>
                <Menu 
                    mode="horizontal"
                    defaultSelectedKeys={["1"]}
                    items={menuItems}
                    style={{ 
                        flex: 1, 
                        minWidth: 0,
                        backgroundColor: "var(--secondary-color)"
                    }}
                    
                />
            </Header>
            <Content style={{ flex: 1, backgroundColor: "var(--primary-color)" }}>
                <Outlet/>
            </Content>
            <Footer style={{ textAlign: "center", backgroundColor: "var(--secondary-color)" }}>
                Jester - a budget management tool - Studio Clue 2024
            </Footer>
        </Layout>
    )
}

export default DashboardLayout;