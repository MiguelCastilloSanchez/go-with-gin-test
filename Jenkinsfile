pipeline {
    agent any

    environment {
        IMAGE_NAME = "sicei"
    }

    stages {
        stage('Delete previous container'){
            steps {
                sh "docker stop sicei-app || true"
                sh "docker rm sicei-app || true"
            }
        }
        stage('Build new image') {
            steps {
                sh "docker build --tag ${IMAGE_NAME}:${BUILD_ID} ."
            }
        }
        stage('Run new container'){
            steps {
                sh "docker run --name sicei-app -p 8081:8080 ${IMAGE_NAME}:${BUILD_ID} -d"
            }
        }
    }
}