import { Button, Card, Col, Row} from 'antd';
import { useForm } from 'react-hook-form';
import TextField from '../components/ui/text-field';
import { useMutation } from "@tanstack/react-query"
import { login } from '../services/authentication';
import { useMessageApi } from '../components/context/message';
import type { MessageInstance } from 'antd/es/message/interface';
import { useNavigate } from 'react-router-dom';

type LoginValues = {
    username: string;
    password: string;
}

const Login = () => {


    //navigation
    const navigate = useNavigate();

    //form stuff
    const { control, handleSubmit, formState: { errors } } = useForm<LoginValues>();
    
    const messageApi = useMessageApi() as MessageInstance
    
    //mutation helpers
    const handleSuccess = () => {
        messageApi.success("Login successful!");
        navigate('/', {replace: true});
    }

    const handleError = () => {
        messageApi.error("Something went wrong. Please try again.");
    }
    
    //mutation
    const { mutate, isPending } = useMutation({
        mutationFn: login,
        onSuccess: handleSuccess,
        onError: handleError,
    });


    //functions
    const onSubmit = (data: LoginValues) => {
        
        const submitData = {
            username: data.username,
            password: data.password,
        }
        
        mutate(submitData);
    }

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
                <form onSubmit={handleSubmit(onSubmit)}>
                    <Row gutter={16}>
                        <Col span={24}>
                            <TextField
                                inputProps={{
                                    placeholder: "Username",
                                }}
                                label="Username"
                                name="username"
                                rules={{ required: "Username is required" }}
                                control={control}
                                error={errors.username?.message as string}
                            />
                        </Col>

                        <Col span={24}>
                            <TextField
                                inputProps={{
                                    placeholder: "Password",
                                }}
                                label="Password"
                                name="password"
                                rules={{ required: "Password is required" }}
                                control={control}
                                error={errors.password?.message as string}
                            />
                        </Col>

                        <Col span={24} style={{ textAlign: "center", marginTop: 16 }}>
                            <Button
                                type="primary"
                                htmlType="submit"
                                style={{ width: '100%' }}
                                loading={isPending}
                            >
                                Login
                            </Button>
                        </Col>
                    </Row>
                </form>
            </Card>
        </div>
    )
}

export default Login;