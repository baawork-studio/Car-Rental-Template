import { Button, Card, CardActions, CardContent, CardMedia, Stack, Typography } from '@mui/material'
import type { Car } from '../types/rental'

type Props = { car: Car; onDetails: (car: Car) => void; onRent: (car: Car) => void }

export function CarRentalCard({ car, onDetails, onRent }: Props) { return <Card><CardMedia component="img" height="130" image={car.image} alt={car.name} /><CardContent><Typography sx={{ fontWeight: 800 }}>{car.name}</Typography><Typography variant="body2" color="text.secondary">เหลือ {car.available} คัน · ฿{car.price.toLocaleString()}/วัน</Typography></CardContent><CardActions><Button size="small" onClick={() => onDetails(car)}>ดูรายละเอียด</Button><Button size="small" variant="contained" onClick={() => onRent(car)}>เช่ารถ</Button></CardActions></Card> }
