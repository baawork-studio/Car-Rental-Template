import { Box, Chip, Stack, Typography } from '@mui/material'
import { CarRentalCard } from '../../components/CarRentalCard'
import { cars, type Car } from '../../types/rental'

type Props = { onDetails: (car: Car) => void; onRent: (car: Car) => void }
export function VehicleCatalogPage({ onDetails, onRent }: Props) { return <Stack spacing={2}><Box><Typography variant="h5" sx={{ fontWeight: 900 }}>เช่ารถขับเอง</Typography><Typography color="text.secondary">เลือกรถที่พร้อมให้บริการใกล้คุณ</Typography></Box><Chip label="รถพร้อมเช่า" color="success" size="small" sx={{ alignSelf: 'start' }} /><Box sx={{ display: 'grid', gridTemplateColumns: 'repeat(2, minmax(0, 1fr))', gap: 1.5 }}>{cars.map((car) => <CarRentalCard key={car.id} car={car} onDetails={onDetails} onRent={onRent} />)}</Box></Stack> }
