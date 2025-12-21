import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { RouterProvider } from 'react-router-dom'
import './App.css'
import { router } from './router'
import { MessageApiProvider } from './components/context/message'
import { AuthProvider } from './components/context/auth'

//query client
const queryClient = new QueryClient()

function App() {

  return (
    <MessageApiProvider>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <RouterProvider router={router}/>
        </AuthProvider>
      </QueryClientProvider>
    </MessageApiProvider>
  )
}

export default App
