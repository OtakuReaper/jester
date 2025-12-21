import { Button, Col, Row, Table, Typography } from "antd";

const Home = () => {

    //budgets
    const budgetData = [
        {
            key: '1',
            budget: 'Pool',
            allocated: '$500',
            spent: '$150',
            remaining: '$350'
        },
        {
            key: '2',
            budget: 'Land Debt',
            allocated: '$300',
            spent: '$200',
            remaining: '$100'
        }
    ]

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