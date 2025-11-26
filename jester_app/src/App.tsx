// import axios from "axios"
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import PrivateLayout from "./components/layouts/private";
import Home from "./modules/home";

const router = createBrowserRouter([
    {
      path: "/",
      element: <PrivateLayout/>,
      children: [
        {
          path: "/",
          element: <Home />,
        }
      ],
    }
  ])

function App() {
  return (
    <RouterProvider router={router}/>
  )
}

export default App
