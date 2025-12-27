import { Button, Col, Row, Table, Typography } from "antd";
import useAuthenticatedQuery from "../hooks/auth-query";
import { getBudgets } from "../services/budgets";
import { useEffect, useState } from "react";
import { set } from "react-hook-form";

type BudgetDisplay = {
    key: string,
    budget: string,
    allocated: string,
    spent: string,
    remaining: string,
}

type Budget = {
    id: string,
    budgetTypeId: string,
    userId: string,
    name: string,
    description: string,
    colour: string,
    allocated: number,
    spent: number,
    amount: number,
}


const Home = () => {
    const userId = "1";

    //states
    const [ budgetData, setBudgetData ] = useState<BudgetDisplay[]>([]);

    // budgets
    const { data: budgets = [], isLoading: budgetIsLoading } = useAuthenticatedQuery<Budget[]>({
        queryKey: ["budgets", userId],
        queryFn: () => getBudgets({ id: userId }),
    });

    useEffect(() => {

        if(!budgetIsLoading && budgets.length != 0 ) {

            const newFormattedBudgets = budgets.map((budget) => {
                return {
                    key: budget.id,
                    budget: budget.name,
                    allocated: `$${budget.allocated.toFixed(2)}`,
                    spent: `$${budget.spent.toFixed(2)}`,
                    remaining: `$${budget.amount.toFixed(2)}`,
                }
            });

            setBudgetData(newFormattedBudgets);
        }
    }, [budgets]);

    const budgetColumns = [
        {
            title: 'Budget',
            dataIndex: 'budget',
            key: 'budget',
        },
        {
            title: 'Allocated',
            dataIndex: 'allocated',
            key: 'allocated',
        },
        {
            title: 'Spent',
            dataIndex: 'spent',
            key: 'spent',
        },
        {
            title: 'Remaining',
            dataIndex: 'remaining',
            key: 'remaining',
        }
    ]

    const entriesData = [
        {
            key: '1',
            description: "Starting Funds",
            date: "2025-12-19",
            amount: "$1209.23",
            budget: "Pool",
            type: "Credit",
        },
        {
            key: '2',
            description: "Land Payment",
            date: "2025-12-20",
            amount: "$500.00",
            budget: "Land Debt",
            type: "Debit",
        }
    ];

    const entriesColumns = [
        {
            title: 'Description',
            dataIndex: 'description',
            key: 'description',
        },
        {
            title: 'Date',
            dataIndex: 'date',
            key: 'date',
        },
        {
            title: 'Amount',
            dataIndex: 'amount',
            key: 'amount',
        },
        {
            title: 'Budget',
            dataIndex: 'budget',
            key: 'budget',
        },
        {
            title: 'Type',
            dataIndex: 'type',
            key: 'type',
        }
    ]


    //TODO: figure out the table footer totals
    
    return (
    <>
        <Row gutter={10} style={{ padding: 20 }}>
            <Col xs={24} md={12} lg={8}>
                <Typography.Title level={4}>Budgets</Typography.Title>
                <Table 
                dataSource={budgetData} 
                columns={budgetColumns} 
                pagination={false}
                size={"small"}
                />
            </Col>

            <Col xs={24} md={12} lg={16}>
                <Typography.Title level={4}>Entries</Typography.Title>
                <Table 
                dataSource={entriesData} 
                columns={entriesColumns} 
                pagination={false}
                size={"small"}
                />
                <Row
                    justify="end"
                >
                    <Button
                        style={{ marginTop: 10 }}
                        type="primary"
                    >
                        Add Entry
                    </Button>
                </Row>
            </Col>
        </Row>
    </>
    );
}

export default Home;