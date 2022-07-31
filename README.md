# Helm + Kubernetes (Minikube) + Golang + Microservice example
In this project I set up four microservices and one API endpoint to calculate numbers.

The project uses `docker-compose` in development mode with hot reload.

And it uses K8s with Helm for staging environment, with a single deploy command.

## Application Layer

```
api ->  addition
    ->  subtraction
    ->  multiplication
    ->  division
```

The api has an endpoint `calculate` which accepts query params like so:

```
/calculate?=&add=2400&add=200&subtract=450&multiply=7&multiply=10&divide=5
```

On sending a `GET` request in this format, it relays the respective calculation to the concerned microservice and responds with the calculated amount by reading the responses internally.

## Infra Layer

### Development Setup
Just do `make dev` to start development environment

Then send a request like so. 

```
curl --request GET \
  --url 'http://localhost:8080/calculate?=&add=2400&add=200&subtract=450&multiply=7&multiply=10&divide=5'
```

Clean up with `make clean`

### Staging Setup
Just do `make staging` to start staging environment with production build

Check if all pods are running:
```
kubectl get po -n calculator
```

Get the local IP Address from Minikube:
```
minikube service list
|-------------|----------------|--------------|---------------------------|
|  NAMESPACE  |      NAME      | TARGET PORT  |            URL            |
|-------------|----------------|--------------|---------------------------|
| calculator  | addition       |         8080 | http://<your-ip>:30010 |
| calculator  | api            |         8080 | http://<your-ip>:30020 |
| calculator  | division       |         8080 | http://<your-ip>:30030 |
| calculator  | multiplication |         8080 | http://<your-ip>:30040 |
| calculator  | subtraction    |         8080 | http://<your-ip>:30050 |
| default     | kubernetes     | No node port |
| kube-system | kube-dns       | No node port |
|-------------|----------------|--------------|---------------------------|

```

Then send a request like so to `api` microservice

```
curl --request GET \
  --url 'http://<your-ip>:30020/calculate?=&add=2400&add=200&subtract=450&multiply=7&multiply=10&divide=5'
```

Clean up with `make stop`

## Requirements
1. docker
2. docker-compose
3. minikube
4. kubectl
5. helm
