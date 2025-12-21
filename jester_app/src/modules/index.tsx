import { Layout, Menu, Typography } from "antd";
import { Content, Footer, Header, } from "antd/es/layout/layout";
import { Outlet } from "react-router-dom"


const DashboardLayout = () => {


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