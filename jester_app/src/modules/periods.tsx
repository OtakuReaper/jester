import { Col, Row } from "antd";

const Periods = () => {
    return (
        <>
            <p>New Period</p>
            <p>Periods are what define the budgeting cycle for a specific timeframe. ... Other Info Text</p>

            <Row gutter={10} style={{ padding: 20 }}>
                <Col xs={24} md={12} lg={8}>
                    <div>Period Table Here</div>
                </Col>
            </Row>

        </>
    )
}

export default Periods;