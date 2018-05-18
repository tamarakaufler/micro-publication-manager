protoc:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tamarakaufler/publication-manager/book-service \
	  proto/book.proto
	
build: protoc
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		docker build -t book-service .

run: 
		docker run -p 50053:50051 \
    -e MICRO_SERVER_ADDRESS=:50051 \
    -e MICRO_REGISTRY=mdns book-service
