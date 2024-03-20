.PHONY: docker-build-api check-api run-api

docker-build-api:
	docker build -t api:development -f deployments/development/Dockerfile.api .

check-api:
	curl localhost:8080/status

run-api:
	docker-compose -f deployments/development/docker-compose.yml up -d

.DEFAULT_GOAL := run-api