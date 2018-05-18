build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o author-cli
	docker build -t author-cli .

run:
	docker run -e MICRO_REGISTRY=kubernetes author-cli