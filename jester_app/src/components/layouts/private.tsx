import { Dropdown, Layout, Row, Space, Typography, type MenuProps } from "antd";
import { Outlet } from "react-router-dom";
import { MenuOutlined } from "@ant-design/icons";

const navItems: MenuProps["items"] = [
    {
        key: '1',
        label: 'Budgets',
    },
    {
        key: '2',
        label: 'Reports',
    },
]

const PrivateLayout: React.FC = () => {
    return (
        <Layout>
            <Layout.Header style={{ padding: 0, background: "var(--secondary-bg)"}}>
                <Row justify={"space-between"} style={{
                    paddingTop: "0.2em",
                    paddingBottom: "0.2em",
                }}>
                    <h1 style={{ 
                        color: "var(--main-text)", 
                        fontSize: "3em", 
                        paddingLeft: "0.2em", 
                        margin: 0, 
                        backgroundColor: "orange"
                    }}>
                        Jester
                    </h1>
                    <div style={{ paddingRight: "0.2em", backgroundColor: "orange" }}>
                        <Dropdown
                            menu={{ items: navItems, selectable: true,}}
                        >
                            <Typography.Text style={{ 
                                color: "var(--main-text)", 
                                fontSize: "2.5em", 
                                cursor: "pointer", 
                                userSelect: "none",
                                margin: 0,
                                paddingInline: "0.25em",
                                }}>
                                <Space>
                                    <MenuOutlined />
                                </Space>
                            </Typography.Text>
                        </Dropdown>
                    </div>
                </Row>
            </Layout.Header>
            <Outlet />
        </Layout>
    )
}

export default PrivateLayout;