package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	publisherProto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
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

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	microSrv := k8s.NewService(
		micro.Name("publication.management.book"),
		micro.Version("latest"),
	)
	pubClient := publisherProto.NewPublisherServiceClient("publication.management.publisher", microSrv.Client())

	microSrv.Init()

	proto.RegisterBookServiceHandler(microSrv.Server(), &service{session, pubClient})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
