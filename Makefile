dev:
	docker-compose build
	docker-compose up

clean:
	@echo "Cleaning environment..."
	docker-compose stop
	docker-compose down -v

release:
	docker build -t donnieashok/api:prod ./api -f ./Dockerfile
	docker build -t donnieashok/addition:prod ./addition -f ./Dockerfile
	docker build -t donnieashok/division:prod ./division -f ./Dockerfile
	docker build -t donnieashok/multiplication:prod ./multiplication -f ./Dockerfile
	docker build -t donnieashok/subtraction:prod ./subtraction -f ./Dockerfile
	docker push donnieashok/api:prod
	docker push donnieashok/addition:prod
	docker push donnieashok/division:prod
	docker push donnieashok/multiplication:prod
	docker push donnieashok/subtraction:prod

staging:
	minikube start
	helm install calculator ./helm
	minikube service api -n calculator

stop:
	@echo "Cleaning environment..."
	helm uninstall calculator
	minikube stop
	minikube delete --all
