PROJECT?=github.com/lallison/h_skills_project
NAME?=api
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

.PHONY: docker-build-api check-api run-api

build-api:
	go build \
	-ldflags "-w -s \
	-X ${PROJECT}/version.Version=${VERSION} \
	-X ${PROJECT}/version.Name=${NAME} \
	-X ${PROJECT}/version.Commit=${COMMIT} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o bin/api cmd/api/main.go 

docker-build-api:
	docker build -t api:development -f deployments/development/Dockerfile.api .

check-api:
	curl localhost:8080/status | jq

run-api:
	docker-compose -f deployments/development/docker-compose.yml up -d

.DEFAULT_GOAL := run-api
