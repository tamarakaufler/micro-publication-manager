build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book-cli
	docker build -t book-cli .

run:
	docker run --net="host" -e MICRO_REGISTRY=mdns book-cli
