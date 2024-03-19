include scripts/*.mk

.PHONY: dev-up dev-down

dev-up:
	docker-compose -f deployments/development/docker-compose.yml up -d

dev-down:
	docker-compose -f deployments/development/docker-compose.yml down -v

.DEFAULT_GOAL := dev-up