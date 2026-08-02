import { Box, TextField } from '@mui/material'

type Props = { pickupDate: string; returnDate: string; onChange: (field: 'pickupDate' | 'returnDate', value: string) => void }

export function DateRangePicker({ pickupDate, returnDate, onChange }: Props) {
  return <Box sx={{ position: 'relative', display: 'grid', gap: 2 }}><Box sx={{ display: { xs: 'none', sm: 'block' }, position: 'absolute', top: 28, left: '30%', right: '30%', borderTop: '2px solid', borderColor: 'primary.light', zIndex: 0 }} /><TextField label="วันรับรถ" type="date" value={pickupDate} onChange={(e) => onChange('pickupDate', e.target.value)} slotProps={{ inputLabel: { shrink: true }, htmlInput: { min: new Date().toISOString().slice(0, 10) } }} /><TextField label="วันคืนรถ" type="date" value={returnDate} disabled={!pickupDate} onChange={(e) => onChange('returnDate', e.target.value)} slotProps={{ inputLabel: { shrink: true }, htmlInput: { min: pickupDate } }} /></Box>
}
