import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { router } from './router' 
import { AuthContextProvider } from './AuthContext'

import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthContextProvider>
      <RouterProvider router={ router }/>
    </AuthContextProvider> 
  </StrictMode>,
)
