# k8s-nginx

This is my playground for kubernetes.

## Deployment and pods

Create a deployment and pods by `kubectl create` command`,

```sh
kubectl create -f deployment.yaml
```

Create a service by `kubectl create` command`,

```sh
kubectl create -f service.yaml
```

Check URL to service for minikube node.

```sh
minikube service nginx-service --url
```

## Trouble shoothing

CHecking deployment:

```sh
kubectl get deployments
kubectl describe deployments 
```

Checking created and pods:

```sh
kubectl get pods
kubectl describe pods
```

Checking services:

```sh
kubectl get services
kubectl describe service
```

Isolate the cause of the problem

```sh
# Get pods' IP address
kubectl get pods --output=wide

# Login to node
minikube ssh

# Check in node
docker ps
curl ${pod's ip}
```

Open kubernetes dashboard.
```sh
minikube service --namespace kube-system  kubernetes-dashboard
```
