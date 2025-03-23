pipeline {
    agent any

    environment {
        IMAGE_NAME = "sicei"
    }

    stages {
        stage('Delete previous container'){
            steps {
                sh "docker stop sicei-app"
                sh "docker rm sicei-app"
            }
        }
        stage('Build new image') {
            steps {
                sh "docker build --tag ${IMAGINE_NAME}:${BUILD_ID} ."
            }
        }
        stage('Run new container'){
            steps {
                sh "docker run --name sicei-app -p 8081:8080 ${IMAGINE_NAME}:${BUILD_ID}"
            }
        }
    }
}