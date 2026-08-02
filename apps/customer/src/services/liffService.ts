import liff from '@line/liff'

export type LineProfile = { userId: string; displayName: string }

export async function initializeLiff(): Promise<LineProfile> {
  const liffId = import.meta.env.VITE_LIFF_ID
  if (!liffId) throw new Error('ไม่พบ LIFF ID')
  await liff.init({ liffId })
  if (!liff.isInClient() || !liff.isLoggedIn()) throw new Error('กรุณาเปิดผ่าน LINE OA เท่านั้น')
  const profile = await liff.getProfile()
  return { userId: profile.userId, displayName: profile.displayName }
}

export async function sendBookingToChat(summary: string) {
  if (liff.isInClient()) await liff.sendMessages([{ type: 'text', text: summary }])
}
