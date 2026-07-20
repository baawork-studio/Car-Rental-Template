import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { CssBaseline, ThemeProvider } from '@mui/material'
import { AdminApp } from './app/AdminApp'
import { adminTheme } from './theme/adminTheme'
import './styles/global.css'

createRoot(document.getElementById('root')!).render(<StrictMode><ThemeProvider theme={adminTheme}><CssBaseline /><AdminApp /></ThemeProvider></StrictMode>)
