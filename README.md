# DIT Library — Plateforme de Gestion de Bibliothèque

Application web fullstack de gestion de bibliothèque académique, conçue avec une architecture microservices et des pratiques DevOps modernes.

## Objectif

Moderniser la gestion des bibliothèques en proposant :

* suivi des livres en temps réel
* gestion des utilisateurs
* système d’emprunt et de retour
* détection automatique des retards

---

## Architecture

* **Backend** : Go (Gin + GORM)
* **Bases de données** : PostgreSQL (1 par microservice)
* **Frontend** : React + Vite
* **Reverse Proxy** : Nginx
* **Conteneurisation** : Docker & Docker Compose
* **CI/CD** : Jenkins

---

## Microservices

| Service              | Port | Description              |
| -------------------- | ---- | ------------------------ |
| service-livres       | 8001 | Gestion des livres       |
| service-utilisateurs | 8002 | Gestion des utilisateurs |
| service-emprunts     | 8003 | Gestion des emprunts     |

---

## Fonctionnalités

* CRUD des livres
* Recherche avancée (titre, auteur, ISBN)
* Gestion des utilisateurs (étudiant, professeur, admin)
* Emprunt et retour de livres
* Détection automatique des retards
* Architecture scalable microservices

---

## Lancer le projet

```bash
git clone https://github.com/Rhupthur/dit-library.git
cd dit-library
docker compose up --build
```

Accès : http://localhost

---

## Structure du projet

```
dit-library/
├── services/
│   ├── livres/
│   ├── utilisateurs/
│   └── emprunts/
├── frontend/
├── nginx/
├── docker-compose.yml
├── Jenkinsfile
└── README.md
```

---

## API

### Livres

| Méthode | Route           |
| ------- | --------------- |
| GET     | /livres         |
| GET     | /livres?titre=x |
| GET     | /livres/:id     |
| POST    | /livres         |
| PATCH   | /livres/:id     |
| DELETE  | /livres/:id     |

---

### Utilisateurs

| Méthode | Route             |
| ------- | ----------------- |
| GET     | /utilisateurs     |
| POST    | /utilisateurs     |
| PATCH   | /utilisateurs/:id |
| DELETE  | /utilisateurs/:id |

---

### Emprunts

| Méthode | Route                    |
| ------- | ------------------------ |
| GET     | /emprunts                |
| POST    | /emprunts                |
| POST    | /emprunts/:id/retour     |
| GET     | /emprunts/historique/:id |

---

## CI/CD

Pipeline Jenkins :

1. Build des images Docker
2. Tests
3. Déploiement automatique
4. Vérification des services

---

## Développement local

```bash
cd services/livres
cp .env.example .env
go run .
```

---

## Améliorations futures

* Authentification (JWT)
* Dashboard frontend
* Monitoring (Prometheus/Grafana)
* Kubernetes

---

## Auteur

Projet DevOps — Dakar Institute of Technology
