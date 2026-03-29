import { useEffect, useState } from 'react'
import { livresAPI } from '../api'

export default function Livres() {
    const [livres, setLivres] = useState([])
    const [form, setForm] = useState({ titre: '', auteur: '', isbn: '' })
    const [recherche, setRecherche] = useState('')
    const [erreur, setErreur] = useState('')

    const chargerLivres = async () => {
        try {
            const params = recherche ? { titre: recherche } : {}
            const res = await livresAPI.get('', { params })
            setLivres(res.data)
        } catch {
            setErreur('Erreur lors du chargement des livres')
        }
    }

    useEffect(() => { chargerLivres() }, [])

    const creerLivre = async (e) => {
        e.preventDefault()
        setErreur('')
        try {
            await livresAPI.post('', form)
            setForm({ titre: '', auteur: '', isbn: '' })
            chargerLivres()
        } catch (err) {
            setErreur(err.response?.data?.error || 'Erreur création')
        }
    }

    const supprimerLivre = async (id) => {
        try {
            await livresAPI.delete(`/${id}`)
            chargerLivres()
        } catch {
            setErreur('Erreur suppression')
        }
    }

    return (
        <div style={styles.container}>
            <h1>Livres</h1>

            <form onSubmit={creerLivre} style={styles.form}>
                <input
                    placeholder="Titre"
                    value={form.titre}
                    onChange={e => setForm({ ...form, titre: e.target.value })}
                    style={styles.input}
                    required
                />
                <input
                    placeholder="Auteur"
                    value={form.auteur}
                    onChange={e => setForm({ ...form, auteur: e.target.value })}
                    style={styles.input}
                    required
                />
                <input
                    placeholder="ISBN"
                    value={form.isbn}
                    onChange={e => setForm({ ...form, isbn: e.target.value })}
                    style={styles.input}
                    required
                />
                <button type="submit" style={styles.btn}>Ajouter</button>
            </form>

            <div style={styles.recherche}>
                <input
                    placeholder="Rechercher par titre..."
                    value={recherche}
                    onChange={e => setRecherche(e.target.value)}
                    style={styles.input}
                />
                <button onClick={chargerLivres} style={styles.btn}>Rechercher</button>
            </div>

            {erreur && <p style={styles.erreur}>{erreur}</p>}

            <table style={styles.table}>
                <thead>
                    <tr>
                        <th>Titre</th>
                        <th>Auteur</th>
                        <th>ISBN</th>
                        <th>Disponible</th>
                        <th>Quantité</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {livres.map(livre => (
                        <tr key={livre.ID}>
                            <td>{livre.titre}</td>
                            <td>{livre.auteur}</td>
                            <td>{livre.isbn}</td>
                            <td>{livre.disponible ? '✅' : '❌'}</td>
                            <td>{livre.quantite}</td>
                            <td>
                                <button
                                    onClick={() => supprimerLivre(livre.ID)}
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
    recherche: { display: 'flex', gap: '1rem', marginBottom: '1rem' },
    input: { padding: '0.5rem', fontSize: '1rem', borderRadius: '4px', border: '1px solid #ccc' },
    btn: { padding: '0.5rem 1rem', backgroundColor: '#1a1a2e', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' },
    btnDanger: { padding: '0.5rem 1rem', backgroundColor: '#e74c3c', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' },
    erreur: { color: 'red' },
    table: { width: '100%', borderCollapse: 'collapse' }
}