import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { router } from './router' 
import { AuthContext } from './AuthContext'

import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthContext.Provider value={ undefined }>
      <RouterProvider router={ router }/>
    </AuthContext.Provider> 
  </StrictMode>,
)
