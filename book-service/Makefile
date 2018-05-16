IMAGE_TAG=alpha

protoc:
	protoc -I/usr/local/include -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tamarakaufler/publication-manager/book-service \
	  proto/book.proto
	
build: protoc
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		docker build -t quay.io/tkaufler/book-service:$(IMAGE_TAG) .

run: 
		docker run -p 50052:50051 \
    -e MICRO_SERVER_ADDRESS=:50051 \
    -e MICRO_REGISTRY=mdns quay.io/tkaufler/book-service:$(IMAGE_TAG)
