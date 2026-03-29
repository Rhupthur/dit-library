import { useEffect, useState } from 'react'
import { empruntsAPI } from '../api'

export default function Emprunts() {
    const [emprunts, setEmprunts] = useState([])
    const [form, setForm] = useState({ livre_id: '', utilisateur_id: '' })
    const [erreur, setErreur] = useState('')

    const chargerEmprunts = async () => {
        try {
            const res = await empruntsAPI.get('')
            setEmprunts(res.data)
        } catch {
            setErreur('Erreur lors du chargement')
        }
    }

    useEffect(() => { chargerEmprunts() }, [])

    const emprunter = async (e) => {
        e.preventDefault()
        setErreur('')
        try {
            await empruntsAPI.post('', {
                livre_id: parseInt(form.livre_id),
                utilisateur_id: parseInt(form.utilisateur_id)
            })
            setForm({ livre_id: '', utilisateur_id: '' })
            chargerEmprunts()
        } catch (err) {
            setErreur(err.response?.data?.error || 'Erreur emprunt')
        }
    }

    const retourner = async (id) => {
        try {
            await empruntsAPI.post(`/${id}/retour`)
            chargerEmprunts()
        } catch (err) {
            setErreur(err.response?.data?.error || 'Erreur retour')
        }
    }

    return (
        <div style={styles.container}>
            <h1>Emprunts</h1>

            <form onSubmit={emprunter} style={styles.form}>
                <input
                    placeholder="ID du livre"
                    type="number"
                    value={form.livre_id}
                    onChange={e => setForm({ ...form, livre_id: e.target.value })}
                    style={styles.input}
                    required
                />
                <input
                    placeholder="ID de l'utilisateur"
                    type="number"
                    value={form.utilisateur_id}
                    onChange={e => setForm({ ...form, utilisateur_id: e.target.value })}
                    style={styles.input}
                    required
                />
                <button type="submit" style={styles.btn}>Emprunter</button>
            </form>

            {erreur && <p style={styles.erreur}>{erreur}</p>}

            <table style={styles.table}>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Livre ID</th>
                        <th>Utilisateur ID</th>
                        <th>Date emprunt</th>
                        <th>Date retour</th>
                        <th>En retard</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {emprunts.map(e => (
                        <tr key={e.ID}>
                            <td>{e.ID}</td>
                            <td>{e.livre_id}</td>
                            <td>{e.utilisateur_id}</td>
                            <td>{new Date(e.date_emprunt).toLocaleDateString()}</td>
                            <td>{e.date_retour ? new Date(e.date_retour).toLocaleDateString() : '—'}</td>
                            <td>{e.en_retard ? '⚠️' : '✅'}</td>
                            <td>
                                {!e.date_retour && (
                                    <button
                                        onClick={() => retourner(e.ID)}
                                        style={styles.btn}
                                    >
                                        Retourner
                                    </button>
                                )}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

const styles = {
    container: { padding: '2rem' },
    form: { display: 'flex', gap: '1rem', marginBottom: '1rem', flexWrap: 'wrap' },
    input: { padding: '0.5rem', fontSize: '1rem', borderRadius: '4px', border: '1px solid #ccc' },
    btn: { padding: '0.5rem 1rem', backgroundColor: '#1a1a2e', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' },
    erreur: { color: 'red' },
    table: { width: '100%', borderCollapse: 'collapse' }
}