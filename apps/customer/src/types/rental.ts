export type Car = { id: string; name: string; available: number; price: number; image: string }
export type BookingDraft = { car: Car; branch: string; name: string; phone: string; pickupDate: string; returnDate: string; pickupTime: string; returnTime: string; extras: string[] }

export const cars: Car[] = [
  { id: 'yaris-ativ', name: 'Toyota Yaris Ativ', available: 3, price: 990, image: 'https://images.unsplash.com/photo-1549317661-bd32c8ce0db2?auto=format&fit=crop&w=600&q=80' },
  { id: 'honda-city', name: 'Honda City', available: 2, price: 1_190, image: 'https://images.unsplash.com/photo-1492144534655-ae79c964c9d7?auto=format&fit=crop&w=600&q=80' },
  { id: 'toyota-fortuner', name: 'Toyota Fortuner', available: 1, price: 2_490, image: 'https://images.unsplash.com/photo-1533473359331-0135ef1b58bf?auto=format&fit=crop&w=600&q=80' },
]
