dev:
	docker-compose build
	docker-compose up

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

kube-manual:
	kubectl create deployment addition --image=donnieashok/addition:prod
	kubectl create deployment api --image=donnieashok/api:prod
	kubectl create deployment division --image=donnieashok/division:prod
	kubectl create deployment multiplication --image=donnieashok/multiplication:prod
	kubectl create deployment subtraction --image=donnieashok/subtraction:prod
	kubectl expose deployment addition --type=LoadBalancer --port=8080
	kubectl expose deployment api --type=LoadBalancer --port=8080
	kubectl expose deployment division --type=LoadBalancer --port=8080
	kubectl expose deployment multiplication --type=LoadBalancer --port=8080
	kubectl expose deployment subtraction --type=LoadBalancer --port=8080
	minikube service api

staging:
	minikube start
	kubectl apply -f k8s

clean:
	@echo "Cleaning Docker environment..."
	docker-compose stop
	docker-compose down -v
	minikube stop
	minikube delete --all
