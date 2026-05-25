pipeline {
    agent any

    environment {
        // Garante que o Go saiba onde encontrar os binários no Jenkins
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
                // O catchError garante que se o teste falhar, o Build fica INSTÁVEL (Amarelo) e não FALHO (Vermelho), cumprindo o Cenário 4.3
                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    sh 'go test -v -coverprofile=coverage.out ./...'
                }
            }
        }

        stage('Metrics (Bônus)') {
            steps {
                echo "Apresentando métricas de cobertura de código..."
                sh 'go tool cover -func=coverage.out'
            }
        }
    }
}
