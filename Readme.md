# Example of an app deployed on Kubernetes

The repo has one Go file app.go which has two endpoints '/' and '/health'

In order to deploy this app onto Kubernetes we need to create a Docker image and push to the Docker Hub

## Steps to push the Docker image of the Go app onto Dockerhub

### Build the Docker Image 
`docker build --tag probes .`

### Spin up a container with the image `probe`
`docker run -p 8080:8080 probe`

### Tag the image so its suits Docker hub format
`docker tag probes mountainpinelake/probe-hello`

### Push the image to Docker Hub
`docker push mountainpinelake/probe-hello:latest`

## Following are the steps to deploy this image onto minikube Kubernetes

1. Start the minikube cluster
    `minikube start`

2. The `deployment.yaml` file specifies the number of pods that are going to be run using this image in addition to many other important details. Use the following command to spin up the pods by applying the deployment

    `kubectl apply -f deployment.yaml`

3. Check if the pods are up and running

    `kubectl get pods`

4. The `service.yaml` file enables you to expose the Pod as a kubernetes Service

    `kubectl apply -f service.yaml`

5. View the Service

    `kubectl get services`

6. Open a browser window that serves your app 

    `minikube service probes-service`

## Debugging

In order to debug your deployment you can use a combination of the following commands

`kubectl get pods -l app=probes`

Get a detailed overview of the Service 

`kubectl describe svc probes-service`

To continously follow logs of a container

`kubectl logs -f <pod_name>`

To Debug a Pod interactively using a shell session
`kubectl exec -it <pod_name> -- /bin/sh`



