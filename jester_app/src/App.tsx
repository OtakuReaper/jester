import axios from "axios"
import './App.css'
import { useEffect, useState } from "react"
import {Row, Table, type TableProps} from "antd"
import Column from "antd/es/table/Column";

function App() {

  const [ping, SetPing] = useState("nothing");

  const instance = axios.create({
    baseURL: "http://localhost:8080",
    timeout: 1000,
    withCredentials: false,
    headers: {
      "Content-Type" : "application/json",
    }
  })


  const pingHandler = async () => {
    const data = await instance.get("/ping");
    SetPing(data.data.message);
  }

  interface budgetType {
    id: Number,
    name: String,
    allocated: Number,
    spent: Number,
    color: String,
  } 

  const budgetColumns: TableProps<budgetType>['columns'] = [
    {
      title: 'Budget',
      dataIndex: 'name',
      key: 'name',
      render: (text) => <p>{text}</p>,
    }
  ]

  const budgets: budgetType[] = [
    { id: 1, name: "Groceries", allocated: 86.60, spent: 86.60, color: "#d9d2e9"},
    { id: 2, name: "Disgressionary Food", allocated: 45, spent: 34.50, color: "#d9d2e9"},
    { id: 3, name: "Gym", allocated: 45, spent: 0, color: "#d9d2e9"},
    { id: 4, name: "Car Gas", allocated: 100, spent: 100, color: "#d9d2e9"}
  ]

  useEffect(() => {
    if (budgets.length > 0) {
      SetPing("Budgets Loaded")
    }
  }, [SetPing, ping])

  return (
    <>  
    <div>
        <p>{budgets[1].name}</p>
    </div>

    <Table<budgetType> dataSource={budgets} columns={budgetColumns} rowKey="id" pagination={false}>

    </Table>
      

      <button onClick={() => pingHandler()}>
        {ping}
      </button>
    </>
  )
}

export default App
