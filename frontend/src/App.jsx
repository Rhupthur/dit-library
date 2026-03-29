import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import Navbar from './components/Navbar'
import Livres from './pages/Livres'
import Utilisateurs from './pages/Utilisateurs'
import Emprunts from './pages/Emprunts'

export default function App() {
    return (
        <BrowserRouter>
            <Navbar />
            <Routes>
                <Route path="/" element={<Navigate to="/livres" />} />
                <Route path="/livres" element={<Livres />} />
                <Route path="/utilisateurs" element={<Utilisateurs />} />
                <Route path="/emprunts" element={<Emprunts />} />
            </Routes>
        </BrowserRouter>
    )
}