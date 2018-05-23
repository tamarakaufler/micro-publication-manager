package main

import (
	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName         = "publication-manager"
	bookCollection = "books"
)

type Datastore interface {
	Create(*proto.Book) error
	GetAll() ([]*proto.Book, error)
	Close()
}

type Store struct {
	session *mgo.Session
}

func (db *Store) Create(book *proto.Book) error {
	return db.collection().Insert(book)
}

// func (db *Store) Get(id string) (*proto.Book, error) {
// 	var book *proto.Book
// 	book.Id = id

// 	if err := db.collection().Find(&book).One(&book); err != nil {
// 		return nil, err
// 	}
// 	return book, nil
// }

func (db *Store) GetAll() ([]*proto.Book, error) {
	var books []*proto.Book

	err := db.collection().Find(nil).All(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (db *Store) Close() {
	db.session.Close()
}

func (db *Store) collection() *mgo.Collection {
	return db.session.DB(dbName).C(bookCollection)
}
