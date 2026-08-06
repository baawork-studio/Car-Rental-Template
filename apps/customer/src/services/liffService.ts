import liff from '@line/liff'
import { runtimeConfig } from '../config/runtimeConfig'

export type LineProfile = { userId: string; displayName: string }

let isLiffReady = false

export async function initializeLiff(): Promise<LineProfile> {
  const liffId = runtimeConfig.liffId
  if (!liffId) throw new Error('ไม่พบ LIFF ID')

  await liff.init({ liffId })
  isLiffReady = true
  if (!liff.isInClient()) throw new Error('กรุณาเปิดผ่าน LINE LIFF เท่านั้น')
  if (!liff.isLoggedIn()) throw new Error('กรุณาเข้าสู่ระบบ LINE ก่อนใช้งาน')

  const profile = await liff.getProfile()
  return { userId: profile.userId, displayName: profile.displayName }
}

export async function sendBookingToChat(summary: string) {
  if (isLiffReady && liff.isInClient()) {
    await liff.sendMessages([{ type: 'text', text: summary }])
  }
}
