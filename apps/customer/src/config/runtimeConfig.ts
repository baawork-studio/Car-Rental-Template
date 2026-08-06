type CustomerRuntimeConfig = {
  apiUrl?: string
  liffId?: string
}

declare global {
  interface Window {
    __CUSTOMER_RUNTIME_CONFIG__?: CustomerRuntimeConfig
  }
}

export const runtimeConfig = window.__CUSTOMER_RUNTIME_CONFIG__ ?? {}
