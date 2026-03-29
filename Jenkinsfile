pipeline {
    agent any

    environment {
        COMPOSE_FILE = 'docker-compose.yml'
    }

    stages {

        stage('Recuperation du code') {
            steps {
                echo 'Clonage du depot GitHub...'
                git branch: 'main', url: 'https://github.com/Rhupthur/dit-library.git'
            }
        }

        stage('Verification de lenvironnement') {
            steps {
                echo 'Verification des outils...'
                sh 'docker --version'
                sh 'docker compose version'
            }
        }

        stage('Build des images Docker') {
            steps {
                echo 'Construction des images...'
                sh 'docker compose build --no-cache'
            }
        }

        stage('Deploiement') {
            steps {
                echo 'Deploiement avec Docker Compose...'
                sh 'docker compose down --remove-orphans || true'
                sh 'docker compose up -d'
            }
        }

        stage('Verification des services') {
            steps {
                echo 'Verification que les services repondent...'
                sh 'sleep 20'
                sh 'docker exec dit-library-service-livres-1 wget -qO- http://localhost:8001/health || exit 1'
                sh 'docker exec dit-library-service-utilisateurs-1 wget -qO- http://localhost:8002/health || exit 1'
                sh 'docker exec dit-library-service-emprunts-1 wget -qO- http://localhost:8003/health || exit 1'
            }
        }
    }

    post {
        success {
            echo 'Deploiement reussi !'
        }
        failure {
            echo 'Echec du pipeline - arret des containers...'
            sh 'docker compose down || true'
        }
    }
}
