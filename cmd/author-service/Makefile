IMAGE_TAG=alpha

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o author-cli
	docker build -t author-cli .
	docker build -t quay.io/tkaufler/book-cli:$(IMAGE_TAG) .

run:
	docker run -e MICRO_REGISTRY=mdns quay.io/tkaufler/book-cli:$(IMAGE_TAG)
