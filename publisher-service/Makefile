IMAGE_TAG=v1alpha1
QUAY_PASS?=biggestsecret

protoc:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tamarakaufler/publication-manager/publisher-service \
    proto/publisher.proto
	
build: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o publisher-service -a -installsuffix cgo main.go \
	datastore.go store.go service.go 
	docker build -t quay.io/tamarakaufler/publisher-service:$(IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/publisher-service:$(IMAGE_TAG)

dev: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o publisher-service -a -installsuffix cgo main.go \
	datastore.go store.go service.go 
	docker build -t quay.io/tamarakaufler/publisher-service:$(IMAGE_TAG) .
	
run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=true \
	publisher-service

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ../deployment/publisher-service-dep.tmpl > ../deployment/book-service-dep.yml
	kubectl apply -f ../deployment/publisher-service-dep.yml	
