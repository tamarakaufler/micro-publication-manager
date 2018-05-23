package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	publisherProto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
	"golang.org/x/net/context"
)

// Service must implement all methods
// defined in the protobuf definition
type service struct {
	session         *mgo.Session
	publisherClient publisherProto.PublisherServiceClient
}

func (s *service) GetRepo() Datastore {
	return &Store{s.session.Clone()}
}

func (s *service) PublishBook(ctx context.Context, req *proto.Book, res *proto.Response) error {
	db := s.GetRepo()
	defer db.Close()

	publisherResponse, err := s.publisherClient.FindAvailable(context.Background(), &publisherProto.Requirement{
		Language: req.Language,
		Copies:   req.Copies,
	})
	log.Printf("Found publisher: %s \n", publisherResponse.Publisher.Name)
	if err != nil {
		return err
	}

	req.PublisherId = publisherResponse.Publisher.Id
	err = db.Create(req)
	if err != nil {
		return err
	}

	res.Registered = true
	res.Book = req
	return nil
}

func (s *service) GetPublishedBooks(ctx context.Context, req *proto.GetRequest, res *proto.Response) error {
	db := s.GetRepo()
	defer db.Close()

	books, err := db.GetAll()
	if err != nil {
		return err
	}
	res.Books = books
	return nil
}
