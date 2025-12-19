import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { RouterProvider } from 'react-router-dom'
import './App.css'
import { router } from './router'
import { MessageApiProvider } from './components/context/message'

//query client
const queryClient = new QueryClient()

function App() {

  return (
    <MessageApiProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router}/>
      </QueryClientProvider>
    </MessageApiProvider>
  )
}

export default App
