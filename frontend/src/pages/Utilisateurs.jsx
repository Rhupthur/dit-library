import { useEffect, useState } from 'react'
import { utilisateursAPI } from '../api'

export default function Utilisateurs() {
    const [utilisateurs, setUtilisateurs] = useState([])
    const [form, setForm] = useState({ nom: '', email: '', type: 'etudiant' })
    const [erreur, setErreur] = useState('')

    const chargerUtilisateurs = async () => {
        try {
            const res = await utilisateursAPI.get('')
            setUtilisateurs(res.data)
        } catch {
            setErreur('Erreur lors du chargement')
        }
    }

    useEffect(() => { chargerUtilisateurs() }, [])

    const creerUtilisateur = async (e) => {
        e.preventDefault()
        setErreur('')
        try {
            await utilisateursAPI.post('', form)
            setForm({ nom: '', email: '', type: 'etudiant' })
            chargerUtilisateurs()
        } catch (err) {
            setErreur(err.response?.data?.error || 'Erreur création')
        }
    }

    const supprimerUtilisateur = async (id) => {
        try {
            await utilisateursAPI.delete(`/${id}`)
            chargerUtilisateurs()
        } catch {
            setErreur('Erreur suppression')
        }
    }

    return (
        <div style={styles.container}>
            <h1>Utilisateurs</h1>

            <form onSubmit={creerUtilisateur} style={styles.form}>
                <input
                    placeholder="Nom"
                    value={form.nom}
                    onChange={e => setForm({ ...form, nom: e.target.value })}
                    style={styles.input}
                    required
                />
                <input
                    placeholder="Email"
                    type="email"
                    value={form.email}
                    onChange={e => setForm({ ...form, email: e.target.value })}
                    style={styles.input}
                    required
                />
                <select
                    value={form.type}
                    onChange={e => setForm({ ...form, type: e.target.value })}
                    style={styles.input}
                >
                    <option value="etudiant">Étudiant</option>
                    <option value="professeur">Professeur</option>
                    <option value="admin">Admin</option>
                </select>
                <button type="submit" style={styles.btn}>Ajouter</button>
            </form>

            {erreur && <p style={styles.erreur}>{erreur}</p>}

            <table style={styles.table}>
                <thead>
                    <tr>
                        <th>Nom</th>
                        <th>Email</th>
                        <th>Type</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {utilisateurs.map(u => (
                        <tr key={u.ID}>
                            <td>{u.nom}</td>
                            <td>{u.email}</td>
                            <td>{u.type}</td>
                            <td>
                                <button
                                    onClick={() => supprimerUtilisateur(u.ID)}
                                    style={styles.btnDanger}
                                >
                                    Supprimer
                                </button>
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
    btnDanger: { padding: '0.5rem 1rem', backgroundColor: '#e74c3c', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' },
    erreur: { color: 'red' },
    table: { width: '100%', borderCollapse: 'collapse' }
}