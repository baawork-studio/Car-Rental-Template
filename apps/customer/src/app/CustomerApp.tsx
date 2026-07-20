import { Box, Button, Container, Stack, Typography } from '@mui/material'
import DirectionsCarRoundedIcon from '@mui/icons-material/DirectionsCarRounded'

export function CustomerApp() {
  return (
    <Box component="main" sx={{ minHeight: '100dvh', display: 'grid', placeItems: 'center', py: 3 }}>
      <Container maxWidth="sm"><Stack spacing={3} alignItems="center" textAlign="center">
        <DirectionsCarRoundedIcon color="primary" sx={{ fontSize: 56 }} />
        <Typography variant="h4" fontWeight={800}>Car Rental</Typography>
        <Typography color="text.secondary">พื้นที่สำหรับลูกค้าบน LINE OA พร้อมเชื่อมต่อ LIFF ในขั้นตอนถัดไป</Typography>
        <Button variant="contained" fullWidth disabled>เริ่มค้นหารถเร็ว ๆ นี้</Button>
      </Stack></Container>
    </Box>
  )
}
