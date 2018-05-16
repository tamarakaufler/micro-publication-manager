IMAGE_TAG=alpha

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book-cli
	docker build -t quay.io/tkaufler/book-cli:$(IMAGE_TAG) .

run:
	docker run -e MICRO_SELECTOR=static quay.io/tkaufler/book-cli:$(IMAGE_TAG)
