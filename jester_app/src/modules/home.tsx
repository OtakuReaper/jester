import { Button, Col, Modal, Row, Table, Typography } from "antd";
import useAuthenticatedQuery from "../hooks/auth-query";
import { getBudgets, getEntries } from "../services/budgets";
import { useEffect, useState } from "react";

type BudgetDisplay = {
    key: string,
    budget: string,
    current_amount: string,
}

type Budget = {
    id: string,
    budgetTypeId: string,
    userId: string,
    name: string,
    description: string,
    colour: string,
    allocation: number,
    spent: number,
    current_amount: number,
}

type Entries = {

}

const Home = () => {
    
    const userId = '4cc85064-a411-4d01-9dda-0adb2d8a1b52'

    //states
    const [ budgetData, setBudgetData ] = useState<BudgetDisplay[]>([]);
    const [ entryData, setEntryData ] = useState<Entries[]>([]);

    const [ openEntryModal, setOpenEntryModal ] = useState<boolean>(false);

    // budgets
    const { data: budgets = [], isLoading: budgetIsLoading } = useAuthenticatedQuery<Budget[]>({
        queryKey: ["budgets", userId],
        queryFn: () => getBudgets({ id: userId }),
    });

    const { data: entries = [], isLoading: entriesIsLoading } = useAuthenticatedQuery<Entries[]>({
        queryKey: ["entries", userId],
        queryFn: () => getEntries({ id: userId}),
    });


    //functions
    const openAddEntryModal = () => {
        setOpenEntryModal(true);
    }

    const handleCancelEntry = () => {
        setOpenEntryModal(false);
    }

    const handleNewEntry = () => {
        //TODO: add mutation logic here
    }

    //rendering
    useEffect(() => {

        if(!budgetIsLoading && budgets.length != 0 ) {

            const newFormattedBudgets = budgets.map((budget) => {

                const curr = budget.current_amount.toFixed(2);

                return {
                    key: budget.id,
                    budget: budget.name,
                    current_amount: `$${curr}`,
                }
            });

            setBudgetData(newFormattedBudgets);
        }

        if(!entriesIsLoading && entries.length != 0 ) {
            setEntryData(entries);    
        }

    }, [budgetIsLoading, entriesIsLoading]);

    const budgetColumns = [
        {
            title: 'Budget',
            dataIndex: 'budget',
            key: 'budget',
        },
        {
            title: 'Remaining',
            dataIndex: 'current_amount',
            key: 'current_amount',
            render: (text: string) => <span style={{color: Number(text.replace('$', '')) <= 0 ? 'red' : 'black' }}>{text}</span>,
        }
    ]

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
        // {
        //     title: 'Type',
        //     dataIndex: 'type',
        //     key: 'type',
        // }
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
                dataSource={entryData} 
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
                        onClick={() => openAddEntryModal()}
                    >
                        Add Entry
                    </Button>
                </Row>
            </Col>

            <Modal
                title="Add Entry"
                open={openEntryModal}
                onOk={handleNewEntry}
                confirmLoading={entriesIsLoading}
                onCancel={handleCancelEntry}
                okText="ADD ENTRY"
                cancelText="CANCEL"
            >
                <p>Entry form goes here</p>
            </Modal>
        </Row>
    </>
    );
}

export default Home;