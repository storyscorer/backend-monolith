.PHONY: build start tidy

build: tidy
	go build .

start: build
	./backend-server

tidy:
	go mod tidy