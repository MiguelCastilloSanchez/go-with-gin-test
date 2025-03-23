pipeline {
    agent any

    environment {
        IMAGE_NAME = "sicei"
        CONTAINER_NAME = "sicei-app"
    }

    stages {
        stage('Delete previous container'){
            steps {
                sh "docker stop ${CONTAINER_NAME} || true"
                sh "docker rm ${CONTAINER_NAME} || true"
            }
        }
        stage('Build new image') {
            steps {
                sh "docker build --tag ${IMAGE_NAME}:${BUILD_ID} ."
            }
        }
        stage('Run new container'){
            steps {
                sh "docker run -d --name ${CONTAINER_NAME} -p 8081:8080 ${IMAGE_NAME}:${BUILD_ID}"
            }
        }
    }
}