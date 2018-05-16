IMAGE_TAG=alpha

protoc:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tamarakaufler/publication-manager/publisher-service \
    proto/publisher.proto
	
build: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t quay.io/tkaufler/publisher-service:$(IMAGE_TAG) .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns quay.io/tkaufler/publisher-service:$(IMAGE_TAG)
