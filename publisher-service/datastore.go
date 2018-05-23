package main

import (
	_ "github.com/micro/go-plugins/registry/kubernetes"

	mgo "gopkg.in/mgo.v2"

	proto "github.com/tamarakaufler/publication-manager/publisher-service/proto"
)

const (
	dbName              = "publication-manager"
	publisherCollection = "publishers"
)

type Datastore interface {
	Create(*proto.Publisher) error
	GetAll() ([]*proto.Publisher, error)
	Close()
}

type Store struct {
	session *mgo.Session
}

func (db *Store) Create(publisher *proto.Publisher) error {
	return db.collection().Insert(publisher)
}

func (db *Store) GetAll() ([]*proto.Publisher, error) {
	var publishers []*proto.Publisher

	err := db.collection().Find(nil).All(&publishers)
	if err != nil {
		return nil, err
	}
	return publishers, err
}

func (db *Store) Close() {
	db.session.Close()
}

func (db *Store) collection() *mgo.Collection {
	return db.session.DB(dbName).C(publisherCollection)
}
