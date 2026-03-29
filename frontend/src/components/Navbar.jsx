import { Link } from 'react-router-dom'

export default function Navbar() {
    return (
        <nav style={styles.nav}>
            <span style={styles.brand}>DIT Library</span>
            <div style={styles.links}>
                <Link to="/livres" style={styles.link}>Livres</Link>
                <Link to="/utilisateurs" style={styles.link}>Utilisateurs</Link>
                <Link to="/emprunts" style={styles.link}>Emprunts</Link>
            </div>
        </nav>
    )
}

const styles = {
    nav: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '1rem 2rem',
        backgroundColor: '#1a1a2e',
        color: 'white'
    },
    brand: {
        fontSize: '1.5rem',
        fontWeight: 'bold'
    },
    links: {
        display: 'flex',
        gap: '2rem'
    },
    link: {
        color: 'white',
        textDecoration: 'none',
        fontSize: '1rem'
    }
}