import axios from 'axios'

const baseDomain = 'https://pokeapi.co/api'
const apiVer = 'v2'
const baseURL = `${baseDomain}/${apiVer}/`

const api = axios.create({ baseURL })

export default api
