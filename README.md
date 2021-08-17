# Microservice example
In this project I set up four microservices and one API endpoint to calculate numbers.

The project uses docker-compose in development mode with hot reload.

The plan is to use kubernetes for production environment, with a single deploy command.

## Structure

```
api ->  addition
    ->  subtraction
    ->  multiplication
    ->  division
```

The api has an endpoint `calculate` which accepts query params like so:

```
http://localhost:8080/calculate?add=400&add=200&subtract=450&multiply=7&multiply=10&divide=5&divide=2
```

On sending a `GET` request in this format, it relays the respective calculation to the concerned microservice and responds with the calculated amount by reading the responses internally.

## Development Setup
Do `make dev` to start development environment

Then send a request like so. 

```
GET http://localhost:8080/calculate?add=400&add=200&subtract=450&multiply=7&multiply=10&divide=5&divide=2
```

Clean up with `make clean`
