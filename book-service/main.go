package main

import (
	"errors"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"
	authorProto "github.com/tamarakaufler/publication-manager/author-service/proto"
	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	publisherProto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
	"golang.org/x/net/context"
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

	// End mongo session
	// before the main function closes.
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	microSrv := k8s.NewService(
		micro.Name("publication.management.book"),
		micro.Version("latest"),
		micro.WrapHandler(Authenticator),
	)
	pubClient := publisherProto.NewPublisherServiceClient("publication.management.publisher", microSrv.Client())

	microSrv.Init()

	proto.RegisterBookServiceHandler(microSrv.Server(), &service{session, pubClient})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Authenticator(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		authorClient := authorProto.NewAuthorServiceClient("publication.manager.author", client.DefaultClient)

		_, err := authorClient.ValidateToken(context.Background(), &authorProto.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
