import { Button, Col, Row, Table, Typography } from "antd";
import { useEffect, useState } from "react";
import useAuthenticatedQuery from "../hooks/auth-query";
import { useAuth } from "../components/context/hook";
import { getPeriods } from "../services/budgets";

type Period = {
    id: string;
    start_date: string;
    end_date: string;
}

type DisplayPeriod = Period & {
    order: number;
}

const Periods = () => {
    
    //hooks
    const { auth } = useAuth();
    const userId = auth?.id as string;

    //states
    const [periodsData, setPeriodsData] = useState<DisplayPeriod[]>([]);

    //periods
    const { data: periods = [], isLoading: periodsIsLoading } = useAuthenticatedQuery<DisplayPeriod[]>({
        queryKey: ["periods", userId],
        queryFn: () => getPeriods({ id: userId }),
        refetchOnWindowFocus: false,
    });

    //rendering
    useEffect(() => {
        if(!periodsIsLoading && periods.length != 0) {
            setPeriodsData(periods);
        }

    }, [periodsIsLoading])

    const periodsColumns = [
        {
            title: '#',
            dataIndex: 'order',
            key: 'order',
        },
        {
            title: 'Start Date',
            dataIndex: 'start_date',
            key: 'start_date',
        },
        {
            title: 'End Date',
            dataIndex: 'end_date',
            key: 'end_date',
        }
    ]
    
    return (
        <>
            <p>New Period</p>
            <p>Periods are what define the budgeting cycle for a specific timeframe. ... Other Info Text</p>
            
            <Row gutter={10} style={{ padding: 20, marginBottom: 0 }}>
                
                <Col span={24}>
                    <Typography.Title level={4}>All Periods</Typography.Title>
                    <Table
                        dataSource={periodsData}
                        columns={periodsColumns}
                        pagination={false}
                        size={"large"}
                    />  
                </Col>
            </Row>

            <Row gutter={10} justify={"end"} style={{ padding: 20 }}>
                <Button type="primary">Create New Period</Button>
            </Row>

        </>
    )
}

export default Periods;