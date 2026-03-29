import axios from 'axios'

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost'

export const livresAPI = axios.create({
    baseURL: `${BASE_URL}/livres`,
    headers: { 'Content-Type': 'application/json' }
})

export const utilisateursAPI = axios.create({
    baseURL: `${BASE_URL}/utilisateurs`,
    headers: { 'Content-Type': 'application/json' }
})

export const empruntsAPI = axios.create({
    baseURL: `${BASE_URL}/emprunts`,
    headers: { 'Content-Type': 'application/json' }
})