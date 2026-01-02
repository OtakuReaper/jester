import { Button, Card, Form, Input} from 'antd';
import { useEffect, useState } from 'react';
import { useAuth } from '../components/context/hook';
import type { Creds } from '../models/auth';
import { isAxiosError } from 'axios';
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import { Link } from 'react-router-dom';

const Login = () => {

    const [ error, setError ] = useState<string | null>(null);
    const { loginHandler, loginError, loginLoading } = useAuth()
    const [submitted, setSubmitted ] = useState<boolean>(false);

    const onFinish = (values: Creds) => {
        setSubmitted(true);
        setError(null);
        loginHandler(values);
    }

    useEffect(() => {

        if(loginError && submitted) {
            if(isAxiosError(loginError)){
                if(loginError.response === undefined) {
                    setError("Network Error. Please try again.");
                    return;
                }
            }

            if (loginError.response.status === 400){
                setError("Invalid username or password.");
                return
            } 
                
            setError("Something went wrong. Please try again.");
        }

    }, [loginError, submitted, setError]);
    
    return (
        <div
            style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '100vh',
                backgroundColor: '#f0f0f0',
            }}
        >
            <Card style={{ width: 350, boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)" }}>
                <div style={{ textAlign: "center", marginBottom: 24}}>
                    <img
                        src="https://media.tenor.com/4EQW6yKEazIAAAAe/jester-jester-lavorre.png"
                        alt="Jester Logo"
                        style={{
                            maxHeight: "100%",
                            maxWidth: "100%",
                            objectFit: "contain"
                        }}
                    />
                </div> 
                <Form
                    name="login_form"
                    initialValues={{ remember: true}}
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="username"
                        rules={[{ required: true, message: "Please enter your username!" }]}
                    >
                        <Input
                            prefix={<UserOutlined className="site-form-item-icon" />}
                            placeholder="Username"
                            autoComplete="username"
                        />
                    </Form.Item>

                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: "Please enter your password!" }]}
                    >
                        <Input
                            prefix={<LockOutlined className="site-form-item-icon" />}
                            type="password"
                            placeholder="Password"
                            autoComplete="current-password"
                        />
                    </Form.Item>

                    <Form.Item>
                        { !loginLoading && 
                            <Link className="login-form-forgot " to="../forgot-password" >
                                Forgot password?
                            </Link> 
                        }
                    </Form.Item>
                    { error && (
                    <Form.Item>
                        <div style={{ color: "red", textAlign: "center" }}>{error}</div>
                    </Form.Item>
                    )}

                    <Form.Item>
                        <Button type="primary" htmlType="submit" style={{ width: "100%" }} loading={loginLoading}>
                        Log in
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    )
}

export default Login;