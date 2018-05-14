IMAGE_TAG:=alpha

protoc:
	protoc -I/usr/local/include -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tamarakaufler/publication-manager/author-service \
	  proto/author.proto
	
build: protoc
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build
		docker build -t author-service .

run: 
		docker run -p 50053:50051 \
    -e MICRO_SERVER_ADDRESS=:50051 \
    -e MICRO_REGISTRY=mdns author-service
