import liff from '@line/liff'

export type LineProfile = { userId: string; displayName: string }

const guestProfile: LineProfile = { userId: '', displayName: '' }
let isLiffReady = false

export async function initializeLiff(): Promise<LineProfile> {
  const liffId = import.meta.env.VITE_LIFF_ID
  if (!liffId) return guestProfile

  try {
    await liff.init({ liffId })
    isLiffReady = true
    if (!liff.isLoggedIn()) return guestProfile

    const profile = await liff.getProfile()
    return { userId: profile.userId, displayName: profile.displayName }
  } catch {
    return guestProfile
  }
}

export async function sendBookingToChat(summary: string) {
  if (isLiffReady && liff.isInClient()) {
    await liff.sendMessages([{ type: 'text', text: summary }])
  }
}
