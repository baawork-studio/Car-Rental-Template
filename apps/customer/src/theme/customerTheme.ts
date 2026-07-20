import { createTheme } from '@mui/material/styles'

export const customerTheme = createTheme({
  palette: { primary: { main: '#2563eb', dark: '#1d4ed8' }, secondary: { main: '#0ea5e9' }, warning: { main: '#f59e0b' }, background: { default: '#f8fafc', paper: '#ffffff' }, text: { primary: '#0f172a', secondary: '#64748b' }, success: { main: '#16a34a' }, error: { main: '#dc2626' } },
  typography: { fontFamily: 'Inter, "Noto Sans Thai", sans-serif' },
  shape: { borderRadius: 12 },
})
