package main

import (
	"errors"

	mgo "gopkg.in/mgo.v2"

	proto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
	"golang.org/x/net/context"
)

// Service must implement all methods
// defined in the protobuf definition
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Datastore {
	return &Store{s.session.Clone()}
}

func (s *service) CreatePublisher(ctx context.Context, req *proto.Publisher, res *proto.Response) error {
	db := s.GetRepo()
	defer db.Close()

	if err := db.Create(req); err != nil {
		return err
	}
	return nil
}

func (s *service) FindAvailable(ctx context.Context, req *proto.Requirement, res *proto.Response) error {
	db := s.GetRepo()
	publishers, err := db.GetAll()
	if err != nil {
		return err
	}

	for _, pub := range publishers {
		if pub.Language != req.Language {
			//log.Printf("\tpublisher lang: %s - book lang: %s\n", pub.Language, req.Language)
			continue
		}
		//log.Printf("\tpublisher : %d - book copies: %d\n", (pub.Capacity - pub.Commitment), req.Copies)
		if req.Copies <= (pub.Capacity - pub.Commitment) {
			res.Publisher = pub
			res.Availability = true
			return nil
		}

	}
	return errors.New("No publisher found by that spec")
}
