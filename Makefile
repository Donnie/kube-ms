dev-compose:
	docker-compose build
	docker-compose up

clean:
	@echo "Cleaning Docker environment..."
	docker-compose stop
	docker-compose down -v

apply:
	kompose convert
	kubectl apply -f addition-service.yaml,api-service.yaml,division-service.yaml,multiplication-service.yaml,subtraction-service.yaml,addition-deployment.yaml,addition-claim0-persistentvolumeclaim.yaml,api-deployment.yaml,api-claim0-persistentvolumeclaim.yaml,division-deployment.yaml,division-claim0-persistentvolumeclaim.yaml,multiplication-deployment.yaml,multiplication-claim0-persistentvolumeclaim.yaml,subtraction-deployment.yaml,subtraction-claim0-persistentvolumeclaim.yaml

dev-minikube:	apply
