pipeline {
    agent any

    environment {
        GO_VERSION = "1.22"  // adjust to your Go version
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                echo "⚙️ Setting up Go environment"
                go version || true
                '''
            }
        }

        stage('Vet') {
            steps {
                sh 'make vet'
            }
        }

        stage('Test') {
            steps {
                sh 'make test'
            }
            post {
                always {
                    junit '**/TEST-*.xml' // optional, if you export test results
                }
            }
        }

        stage('Build') {
            steps {
                sh 'make build'
            }
        }

        stage('Install') {
            steps {
                sh 'make install'
            }
        }

        stage('Package (Optional Docker)') {
            when {
                expression { fileExists('Dockerfile') }
            }
            steps {
                sh 'make docker'
            }
        }
    }

    post {
        success {
            echo "✅ Build and tests successful!"
        }
        failure {
            echo "❌ Build or tests failed!"
        }
    }
}
