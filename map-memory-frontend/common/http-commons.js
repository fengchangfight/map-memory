import axios from 'axios'

export const AXIOS = axios.create({
  baseURL: 'http://localhost:8042',
  withCredentials: true
})
