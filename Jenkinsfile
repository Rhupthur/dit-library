pipeline {
    agent any

    environment {
        COMPOSE_FILE = 'docker-compose.yml'
    }

    stages {

        stage('Récupération du code') {
            steps {
                echo 'Clonage du dépôt GitHub...'
                checkout scm
            }
        }

        stage('Vérification de lenvironnement') {
            steps {
                echo 'Vérification des outils...'
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

        stage('Déploiement') {
            steps {
                echo 'Déploiement avec Docker Compose...'
                sh 'docker compose down --remove-orphans || true'
                sh 'docker compose up -d'
            }
        }

        stage('Vérification des services') {
            steps {
                echo 'Vérification que les services répondent...'
                sh 'sleep 15'
                sh 'curl -f http://localhost/livres || exit 1'
                sh 'curl -f http://localhost/utilisateurs || exit 1'
                sh 'curl -f http://localhost/emprunts || exit 1'
            }
        }
    }

    post {
        success {
            echo 'Déploiement réussi !'
        }
        failure {
            echo 'Échec du pipeline — arrêt des containers...'
            sh 'docker compose down || true'
        }
    }
}