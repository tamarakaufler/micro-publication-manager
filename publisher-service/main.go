package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	proto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
)

type Datastore interface {
	FindAvailable(*proto.Requirement) (*proto.Publisher, error)
}

type Store struct {
	publishers []*proto.Publisher
	//publications map[string][]*bookProto.Book
}

// FindAvailable - checks a requirement against a map of publishers,
func (store *Store) FindAvailable(r *proto.Requirement) (*proto.Publisher, error) {
	for _, pub := range store.publishers {
		if pub.Language != r.Language {
			//log.Printf("\tpublisher lang: %s - book lang: %s\n", pub.Language, r.Language)
			continue
		}

		//log.Printf("\tpublisher : %d - book copies: %d\n", (pub.Capacity - pub.Commitment), r.Copies)
		if r.Copies <= (pub.Capacity - pub.Commitment) {
			return pub, nil
		}
	}
	return nil, errors.New("No publisher found by that spec")
}

// grpc service handler
type service struct {
	store Datastore
}

func (s *service) FindAvailable(ctx context.Context, req *proto.Requirement, res *proto.Response) error {

	// Find the next available publisher
	publisher, err := s.store.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the publisher as part of the response message type
	res.Publisher = publisher
	return nil
}

func main() {
	publishers := []*proto.Publisher{
		&proto.Publisher{Id: "publisher001", Name: "Mannings", Country: "USA", Language: "English", Category: map[string]bool{"Autobiography": true, "Fiction": true, "Programming": true}, Capacity: 50000, Commitment: 45000},
		&proto.Publisher{Id: "publisher002", Name: "LeMonde", Country: "France", Language: "French", Category: map[string]bool{"Fiction": true, "Autobiography": true}, Capacity: 50000, Commitment: 40000},
		&proto.Publisher{Id: "publisher003", Name: "O'Reilly", Country: "USA", Language: "English", Category: map[string]bool{"Science": true, "Autobiography": true}, Capacity: 50000, Commitment: 20000},
	}
	store := &Store{publishers}

	srv := micro.NewService(
		micro.Name("publication.management.publisher"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterPublisherServiceHandler(srv.Server(), &service{store})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
