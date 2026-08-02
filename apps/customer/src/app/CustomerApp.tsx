import { useEffect, useState } from 'react'
import { Alert, AppBar, Box, CircularProgress, Container, IconButton, Toolbar, Typography } from '@mui/material'
import HomeRoundedIcon from '@mui/icons-material/HomeRounded'
import { initializeLiff, sendBookingToChat } from '../services/liffService'
import { BookingFormPage } from '../features/booking/BookingFormPage'
import { PaymentPage } from '../features/payment/PaymentPage'
import { PaymentStatusPage } from '../features/payment/PaymentStatusPage'
import { VehicleCatalogPage } from '../features/vehicleCatalog/VehicleCatalogPage'
import { VehicleDetailsPage } from '../features/vehicleCatalog/VehicleDetailsPage'
import { cars, type BookingDraft, type Car } from '../types/rental'

type Screen = 'catalog' | 'details' | 'booking' | 'payment' | 'pending' | 'verified'
const initialDraft = (car: Car, name: string): BookingDraft => ({ car, branch: 'สาขาสนามบินสุวรรณภูมิ', name, phone: '', pickupDate: '', returnDate: '', pickupTime: '10:00', returnTime: '10:00', extras: [] })

export function CustomerApp() {
  const [screen, setScreen] = useState<Screen>('catalog')
  const [name, setName] = useState('')
  const [selected, setSelected] = useState<Car>(cars[0])
  const [draft, setDraft] = useState<BookingDraft>(initialDraft(cars[0], ''))
  const [error, setError] = useState('')
  const [ready, setReady] = useState(false)
  useEffect(() => { initializeLiff().then((profile) => { setName(profile.displayName); setDraft(initialDraft(cars[0], profile.displayName)); setReady(true) }).catch((reason: Error) => setError(reason.message)) }, [])
  const rent = (car: Car) => { setSelected(car); setDraft(initialDraft(car, name)); setScreen('booking') }
  const chat = async () => { await sendBookingToChat(`ขอเช่ารถ ${draft.car.name}\nรับรถ: ${draft.pickupDate} ${draft.pickupTime}\nคืนรถ: ${draft.returnDate} ${draft.returnTime}`); setScreen('pending') }
  if (error) return <Container sx={{ py: 6 }}><Alert severity="warning">{error}</Alert></Container>
  if (!ready) return <Box sx={{ minHeight: '100dvh', display: 'grid', placeItems: 'center' }}><CircularProgress /></Box>
  return <Box sx={{ minHeight: '100dvh' }}><AppBar position="sticky" elevation={0}><Toolbar><IconButton color="inherit" onClick={() => setScreen('catalog')}><HomeRoundedIcon /></IconButton><Typography sx={{ fontWeight: 900 }}>เช่ารถขับเอง</Typography></Toolbar></AppBar><Container maxWidth="sm" sx={{ py: 2.5 }}>
    {screen === 'catalog' && <VehicleCatalogPage onDetails={(car) => { setSelected(car); setScreen('details') }} onRent={rent} />}
    {screen === 'details' && <VehicleDetailsPage car={selected} onBack={() => setScreen('catalog')} onRent={() => rent(selected)} />}
    {screen === 'booking' && <BookingFormPage draft={draft} onChange={setDraft} onChat={chat} onPayment={() => setScreen('payment')} />}
    {screen === 'payment' && <PaymentPage draft={draft} onCancel={() => setScreen('booking')} onConfirm={() => setScreen('pending')} />}
    {screen === 'pending' && <PaymentStatusPage verified={false} onClose={() => setScreen('catalog')} />}
    {screen === 'verified' && <PaymentStatusPage verified onClose={() => setScreen('catalog')} />}
  </Container></Box>
}
