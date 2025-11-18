import axios from "axios"
import './App.css'
import { useEffect, useState } from "react"

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

  useEffect(() => {

  }, [SetPing, ping])


  return (
    <>  
      <button onClick={() => pingHandler()}>
        {ping}
      </button>
    </>
  )
}

export default App
