import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { CssBaseline, ThemeProvider } from '@mui/material'
import { CustomerApp } from './app/CustomerApp'
import { customerTheme } from './theme/customerTheme'
import './styles/global.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode><ThemeProvider theme={customerTheme}><CssBaseline /><CustomerApp /></ThemeProvider></StrictMode>,
)
