package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	microSrv := k8s.NewService(
		micro.Name("publication.management.publisher"),
		micro.Version("latest"),
	)

	microSrv.Init()

	proto.RegisterPublisherServiceHandler(microSrv.Server(), &service{session})

	if err := microSrv.Run(); err != nil {
		fmt.Println(err)
	}
}
