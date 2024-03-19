.PHONY: docker-build-api check-api run-api

docker-build-api:
	docker build -t api:development -f deployments/development/Dockerfile.api .

check-api:
	curl localhost:3000/status | jq

run-api:
	docker-compose -f depoyments/development/docker-compose.yml up -d

.DEFAULT_GOAL := run-api