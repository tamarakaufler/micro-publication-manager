package main

import (
	"log"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	publisherProto "github.com/tamarakaufler/publication-manager/publisher-service/proto"

	"golang.org/x/net/context"
)

type Datastore interface {
	Add(*proto.Book) error
	GetAll() ([]*proto.Book, error)
}

type Store struct {
	books []*proto.Book
}

func (db *Store) Add(book *proto.Book) error {
	books := append(db.books, book)
	db.books = books
	return nil
}

func (db *Store) GetAll() ([]*proto.Book, error) {
	return db.books, nil
}

// service implements all of the methods to satisfy the service
// defined in our protobuf definition
type service struct {
	database        Datastore
	publisherClient publisherProto.PublisherServiceClient
}

// AddBook - service method to store the book in the database
func (s *service) PublishBook(ctx context.Context, book *proto.Book, res *proto.Response) error {
	publisherResponse, err := s.publisherClient.FindAvailable(context.Background(), &publisherProto.Requirement{
		Language: book.GetLanguage(),
		Category: book.GetCategory(),
		Copies:   book.GetCopies(),
	})
	if err != nil {
		log.Println("No publisher found")
		return err
	}
	log.Printf("Found publisher: %s \n", publisherResponse.Publisher.Name)

	err = s.database.Add(book)
	if err != nil {
		return err
	}

	book.PublisherId = publisherResponse.Publisher.Id
	res.Book = book
	res.Registered = true

	return nil
}

func (s *service) GetPublishedBooks(ctx context.Context, req *proto.GetRequest, res *proto.Response) error {
	books, err := s.database.GetAll()
	if err != nil {
		return err
	}

	res.Books = books

	return nil
}

func main() {

	microSrv := micro.NewService(
		micro.Name("publication.management.book"),
		micro.Version("latest"),
	)

	db := &Store{}
	pubClient := publisherProto.NewPublisherServiceClient("publication.management.publisher", microSrv.Client())

	microSrv.Init()

	proto.RegisterBookServiceHandler(microSrv.Server(), &service{db, pubClient})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
