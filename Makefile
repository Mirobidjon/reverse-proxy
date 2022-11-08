CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

-include .env

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --platform=linux/amd64 --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

run:
	go run cmd/main.go

linter:
	golangci-lint run

push:
	make build-image TAG=1 SERVICE_NAME=reverse-proxy PROJECT_NAME=learn-cloud-0809 REGISTRY=us.gcr.io ENV_TAG=latest 
	make push-image TAG=1 SERVICE_NAME=reverse-proxy PROJECT_NAME=learn-cloud-0809 REGISTRY=us.gcr.io ENV_TAG=latest 
