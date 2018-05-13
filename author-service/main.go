package main

import (
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/tamarakaufler/publication-manager/author-service/proto"

	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type Datastore interface {
	Create(*pb.Author) error
	GetAll() ([]*pb.Author, error)
}

// Store - Dummy databasesitory, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Store struct {
	authors []*pb.Author
}

func (database *Store) Create(author *pb.Author) error {
	authors := append(database.authors, author)
	database.authors = authors
	return nil
}

func (database *Store) GetAll() ([]*pb.Author, error) {
	return database.authors, nil
}

// service implements all of the methods to satisfy the service
// defined in our protobuf definition
type service struct {
	database Datastore
}

// CreateAuthor - service method to store the author in the database
func (s *service) CreateAuthor(ctx context.Context, author *pb.Author, res *pb.Response) error {

	err := s.database.Create(author)
	if err != nil {
		return err
	}

	res.Author = author
	res.Created = true

	return nil
}

func (s *service) GetAuthors(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	authors, err := s.database.GetAll()
	if err != nil {
		return err
	}

	res.Authors = authors

	return nil
}

func main() {

	database := &Store{}

	microSrv := micro.NewService(
		// matches the package name given in the protobuf definition
		micro.Name("publication.management.author"),
		micro.Version("latest"),
	)
	microSrv.Init()

	pb.RegisterAuthorServiceHandler(microSrv.Server(), &service{database})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
