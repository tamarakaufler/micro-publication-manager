package main

import (
	"log"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	proto "github.com/tamarakaufler/publication-manager/author-service/proto"

	"golang.org/x/net/context"
)

type Datastore interface {
	Create(*proto.Author) error
	GetAll() ([]*proto.Author, error)
}

type Store struct {
	authors []*proto.Author
}

func (database *Store) Create(author *proto.Author) error {
	authors := append(database.authors, author)
	database.authors = authors
	return nil
}

func (database *Store) GetAll() ([]*proto.Author, error) {
	return database.authors, nil
}

// service implements all of the methods to satisfy the service
// defined in our protobuf definition
type service struct {
	database Datastore
}

// CreateAuthor - service method to store the author in the database
func (s *service) CreateAuthor(ctx context.Context, author *proto.Author, res *proto.Response) error {
	err := s.database.Create(author)
	if err != nil {
		return err
	}

	res.Created = true
	res.Author = author

	return nil
}

func (s *service) GetAuthors(ctx context.Context, req *proto.GetRequest, res *proto.Response) error {
	authors, err := s.database.GetAll()
	if err != nil {
		return err
	}

	res.Authors = authors

	return nil
}

func main() {

	microSrv := micro.NewService(
		// matches the package name given in the protobuf definition
		micro.Name("publication.management.author"),
		micro.Version("latest"),
	)
	database := &Store{}
	microSrv.Init()

	proto.RegisterAuthorServiceHandler(microSrv.Server(), &service{database})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
