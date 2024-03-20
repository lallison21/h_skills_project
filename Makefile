include scripts/*.mk

.PHONY: dev-up dev-down

dev-up: docker-build-api run-api

dev-down:
	docker-compose -f deployments/development/docker-compose.yml down -v

.DEFAULT_GOAL := dev-up