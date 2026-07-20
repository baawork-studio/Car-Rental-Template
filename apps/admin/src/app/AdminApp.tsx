import { AppBar, Box, Container, Paper, Stack, Toolbar, Typography } from '@mui/material'
import DashboardRoundedIcon from '@mui/icons-material/DashboardRounded'

export function AdminApp() {
  return (<Box minHeight="100dvh"><AppBar position="static"><Toolbar><DashboardRoundedIcon sx={{ mr: 1 }} /><Typography fontWeight={700}>Car Rental Admin</Typography></Toolbar></AppBar><Container maxWidth="xl" sx={{ py: { xs: 2, md: 4 } }}><Paper sx={{ p: { xs: 3, md: 5 } }}><Stack spacing={1}><Typography variant="h4" fontWeight={800}>Dashboard พร้อมใช้งาน</Typography><Typography color="text.secondary">โครงสร้าง responsive สำหรับจัดการรถ การจอง ลูกค้า และ LINE OA จะถูกเพิ่มในขั้นตอนพัฒนาถัดไป</Typography></Stack></Paper></Container></Box>)
}
