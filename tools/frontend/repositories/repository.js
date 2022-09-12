import axios from 'axios'

const baseDomain = process.env.apiUri
const baseURL = `${baseDomain}`

const api = axios.create({ baseURL })

export default api
