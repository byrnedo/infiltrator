.PHONY: build docker

build:
	mkdir -p build
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static"' -v -o build/infiltrator

docker: build
	docker build -t infiltrator .

