pipeline {
    agent any

    environment {
        AZURE_CLIENT_ID = credentials('azure-credentials') // Service Principal App ID
        AZURE_CLIENT_SECRET = credentials('azure-credentials') // Service Principal Password
        AZURE_TENANT_ID = 'af73384e-f91b-4705-b1c6-05fed392027d'
        ACR_NAME = 'golan'
        ACR_REPO_NAME = 'golan-app'
        IMAGE_TAG = 'latest'
        RESOURCE_GROUP = 'dev-opps'
        AKS_CLUSTER_NAME = 'golanapp'
        REPO_URL = 'https://github.com/Mujadid13/golan-app.git'
        HELM_RELEASE_NAME_GOLAN = 'golan-app'
    }

    stages {
        stage('Checkout Code') {
            steps {
                script {
                    git branch: 'main', url: "${REPO_URL}"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $ACR_NAME.azurecr.io/$ACR_REPO_NAME:$IMAGE_TAG -f app/Dockerfile .'
                }
            }
        }

        stage('Login to ACR') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'azure-credentials', usernameVariable: 'AZURE_CLIENT_ID', passwordVariable: 'AZURE_CLIENT_SECRET')]) {
                        sh 'az login --service-principal -u $AZURE_CLIENT_ID -p $AZURE_CLIENT_SECRET --tenant $AZURE_TENANT_ID'
                        sh 'az acr login --name $ACR_NAME'
                    }
                }
            }
        }

        stage('Push Docker Image to ACR') {
            steps {
                script {
                    sh 'docker push $ACR_NAME.azurecr.io/$ACR_REPO_NAME:$IMAGE_TAG'
                }
            }
        }

        stage('Configure kubectl') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'azure-credentials', usernameVariable: 'AZURE_CLIENT_ID', passwordVariable: 'AZURE_CLIENT_SECRET')]) {
                        sh 'az login --service-principal -u $AZURE_CLIENT_ID -p $AZURE_CLIENT_SECRET --tenant $AZURE_TENANT_ID'
                        sh 'az aks get-credentials --resource-group $RESOURCE_GROUP --name $AKS_CLUSTER_NAME --overwrite-existing'
                    }
                }
            }
        }

        stage('Deploy Golan App with Helm') {
            steps {
                script {
                    sh 'helm upgrade --install golan-app helm --set image.repository=$ACR_NAME.azurecr.io/$ACR_REPO_NAME --set image.tag=$IMAGE_TAG --values helm/values.yaml'
                }
            }
        }

        stage('Debug') {
            steps {
                script {
                    sh 'kubectl get all'
                    sh 'kubectl get events'
                    sh 'kubectl describe pods'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
