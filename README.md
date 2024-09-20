# Project Documentation
## Overview
This project involves deploying applications to Azure Kubernetes Service (AKS) using Helm charts, managed with Terraform for infrastructure provisioning, and CI/CD pipelines via Jenkins for streamlined deployment processes.

## Table of Contents
- [Prerequisites](https://github.com/Mujadid13/golan-app2/blob/main/README.md#prerequisites)
* [Terraform Setup](https://github.com/Mujadid13/golan-app2/main/README.md#terraform-setup)
+ [Helm Setup](https://github.com/Mujadid13/golan-app2/main/README.md#helm-setup)
- [Local Setup for Helm Chart](https://github.com/Mujadid13/golan-app2/main/README.md#local-setup-for-helm-chart)
* [Jenkins CI/CD Pipelines](https://github.com/Mujadid13/golan-app2/main/README.md#jenkins-cicd-pipelines)
+ [Additional Information](https://github.com/Mujadid13/golan-app2/main/README.md#additional-information)

## Prerequisites
Before you begin, ensure you have the following installed and configured:

- Terraform: For infrastructure provisioning.
* Helm: For managing Kubernetes applications.
+ Azure CLI: For interacting with Azure resources.
- Jenkins: For managing CI/CD pipelines.
* Docker: For containerizing applications.
+ Kubectl: For interacting with Kubernetes clusters.
## Terraform Setup
 **Configuration Files**: Ensure your Terraform files are properly configured. They should define resources for Azure Container Registry (ACR) and Azure Kubernetes Service (AKS).

#### Initialize Terraform:
- Run `terraform init` to initialize the working directory containing Terraform configuration files.

#### Validate Configuration:
- Execute `terraform validate` to check the syntax and validity of your Terraform files before applying changes.

#### Review the Plan:
- Execute `terraform plan` to review the proposed changes to your infrastructure.

#### Apply the Configuration:
- Apply the configuration with `terraform apply` and confirm the changes when prompted.

#### Verify Deployment:
- Check the Azure portal to confirm that the resources have been created and that ACR is attached to the AKS cluster as expected.

## Helm Setup
Helm Charts: Ensure your Helm charts are configured correctly for your applications, such as `golan-app` and `redis.`

#### Creating a Helm Chart

- **Create a New Chart**: Use the helm create command to generate a new chart with a default structure.
* **Understand the Directory Structure**: Familiarize yourself with the chart’s components, including `Chart.yaml` (metadata), `values.yaml` (default configuration), and the `templates/` folder (Kubernetes manifest templates).
+ **Configure `values.yaml`**: Open the `values.yaml` file and modify the configuration values to suit your application’s needs, such as image repository, tags, and service settings.
- **Modify Templates**: Edit the files in the `templates/` directory to customize the Kubernetes resources based on the values specified in `values.yaml`.

#### Add Helm Repositories:
- Add any necessary Helm repositories if they are not already configured.

#### Install/Upgrade Helm Releases:
- Use Helm commands to install or upgrade your applications based on the configurations in your `values.yaml files`.

#### Verify Deployment:
- Check the status of your Helm releases to ensure the applications are running correctly in your AKS cluster.

## Local Setup for Helm Chart

#### Create a Local Directory for the Helm Repository:
- mkdir -p ~/my-helm-repo

#### Navigate to Your Helm Chart Directory: 
- Go to the directory where your Helm chart (golan-app) is located.

#### Package the Helm Chart:
- helm package .

#### Move the Packaged Chart to the Local Repository:
- mv golan-app-0.1.0.tgz ~/my-helm-repo/

#### Generate the Index File: 
- Navigate to the local repository directory and generate the index:
* cd ~/my-helm-repo
+ helm repo index .

#### Add the Helm Repository: 
- Use the following command to add your local Helm repo:
* helm repo add golan-app http://localhost:8000

#### Open Port 8000 on Your NSG: 
- Ensure that port 8000 is open on your Network Security Group (NSG) to allow access.

#### Update the Helm Repository: 
- Run the following command to update your Helm repo:
* helm repo update

#### Accessing the Repository: 
- Instead of using localhost:8000, use your VM's public IP address to access the Helm repository.

### Notes
- Make sure you have a local HTTP server running on port 8000 to serve the Helm charts.
* You can use a simple server like Python’s built-in HTTP server:
+ cd ~/my-helm-repo
- nohup python3 -m http.server 8000 &

## Jenkins CI/CD Pipelines
#### Overview:
Jenkins pipelines are used to automate the processes of building Docker images, pushing them to Azure Container Registry (ACR), and deploying them to Azure Kubernetes Service (AKS) using Helm.

#### Pipeline Stages:

- **Checkout Code**: Pull the latest code from your repository where Helm charts and Dockerfiles are stored.

* **Build Docker Image**: Build Docker images from your Dockerfile and push them to ACR.

+ **Deploy to AKS**: Deploy the Docker images to AKS using Helm charts. The pipeline should handle upgrading or installing Helm releases with the latest Docker image tags.

#### Configuration Details:

- **Environment Variables**: Configure environment variables for Docker credentials, ACR, Helm chart path, and Kubernetes configuration.

* **Plugins**: Ensure Jenkins is equipped with the necessary plugins for Kubernetes CLI, Docker Pipeline, and Helm.

#### Running Pipelines:
Trigger the Jenkins pipelines to execute the build and deployment processes. Monitor the pipeline runs to ensure they complete successfully and verify the deployment in your AKS cluster.

## Additional Information
#### Troubleshooting:
Consult Jenkins logs, AKS logs, and Terraform output for any issues that arise during setup or deployment.

#### Best Practices:
Follow security and resource management best practices to maintain a reliable deployment process.
