import axios from 'axios'
import { runtimeConfig } from '../config/runtimeConfig'

export const customerApiClient = axios.create({
  baseURL: runtimeConfig.apiUrl || 'http://localhost:8080/api/v1',
  timeout: 10_000,
})
