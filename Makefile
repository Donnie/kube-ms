dev:
	docker-compose build
	docker-compose up

clean:
	@echo "Cleaning Docker environment..."
	docker-compose stop
	docker-compose down -v
