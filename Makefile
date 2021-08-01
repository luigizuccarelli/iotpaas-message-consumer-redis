.PHONY: all clean verify test build container push 

VERSION ?= "1.16.6"

CONTAINER_ENGINE ?= podman

AUTH ?= "--authfile=~/.docker/podman.json"

REGISTRY ?= quay.io/luigizuccarelli

PROJECT ?= iotpaas-message-consumer

all: clean verify test build container push

build: 
	mkdir -p build
	go build -o build -tags real ./...

verify:
	go fmt ./...
	go vet -tags real ./...

test:
	go test -v -tags fake -coverprofile=tests/results/cover.out ./...

cover:
	go tool cover -html=tests/results/cover.out -o tests/results/cover.html

clean:
	rm -rf build/microservice
	go clean ./...

container: 
	$(CONTAINER_ENGINE) build -t $(REGISTRY)/$(PROJECT):$(VERSION) .

push:
	$(CONTAINER_ENGINE) push $(AUTH) $(REGISTRY)/$(PROJECT):$(VERSION) 
