pipeline {
    agent any

    environment {
        IMAGE_NAME = "sicei"
    }

    stages {
        stage('Checkout'){
            steps {
                echo "list all items"
                sh "ls -la"
            }
        }
        stage('Build') {
            steps {
                echo "Building"
            }
        }
        stage('Run'){
            steps {
                echo "Running..."
            }
        }
    }
}