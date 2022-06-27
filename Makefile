.PHONY: build

build:
	go build .

start: build
	./backend-server