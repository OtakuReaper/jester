import { createContext, useContext, type ReactElement } from 'react';
import type { MessageInstance } from 'antd/es/message/interface';
import { message } from 'antd';

const MessageApiContext = createContext<MessageInstance | null>(null);

export const useMessageApi = () => {
    return useContext(MessageApiContext);
}

export const MessageApiProvider = ({children}: {children: ReactElement}) => {
    const [messageApi, contextHolder] = message.useMessage();

    return (
        <MessageApiContext.Provider value={messageApi}>
            {contextHolder}
            {children}
        </MessageApiContext.Provider>
    )
}