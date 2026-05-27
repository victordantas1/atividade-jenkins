pipeline {
    agent any

    triggers {
        cron('* * * * *')
    }

    environment {
        PATH = "${env.PATH}:/usr/local/go/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                echo "Compilando os fontes em Go..."
                sh 'go build -v ./...'
            }
        }

        stage('Test & Coverage') {
            steps {
                echo "Executando casos de teste e cobertura..."
                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    sh 'go test -v -coverprofile=coverage.out ./...'
                }
            }
        }

        stage('Metrics') {
            steps {
                echo "Apresentando métricas de cobertura de código..."
                sh 'go tool cover -func=coverage.out'
            }
        }
    }
}
