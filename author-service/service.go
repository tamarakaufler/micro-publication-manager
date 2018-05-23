package main

import (
	"errors"

	proto "github.com/tamarakaufler/publication-manager/author-service/proto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var (
	ErrAuthorNotFound = errors.New("Author not found")
	ErrInvalidAuthor  = errors.New("Author not valid")
)

type service struct {
	db           *Store
	tokenService TokenService
}

// CreateAuthor hashes the plaintext password and inserts the author in the database
func (s *service) CreateAuthor(ctx context.Context, req *proto.Author, res *proto.Response) error {
	_, ok := s.getAuthorByEmail(ctx, req)
	if ok {
		return ErrAuthorExists
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hp)
	if err := s.db.Create(req); err != nil {
		return err
	}
	res.Created = true
	return nil
}

func (s *service) getAuthorByEmail(ctx context.Context, req *proto.Author) (*proto.Author, bool) {
	author, err := s.db.GetByEmail(req.Email)
	if err != nil {
		return author, true
	}
	return nil, false
}

func (s *service) GetAuthor(ctx context.Context, req *proto.Author, res *proto.Response) error {
	author, err := s.db.Get(req.Id)
	if err != nil {
		return ErrAuthorNotFound
	}
	res.Author = author
	return nil
}

func (s *service) GetAll(ctx context.Context, req *proto.GetAllRequest, res *proto.Response) error {
	authors, err := s.db.GetAll()
	if err != nil {
		return err
	}
	res.Authors = authors
	return nil
}

func (s *service) Authenticate(ctx context.Context, req *proto.Author, res *proto.Token) error {
	author, err := s.db.GetByEmail(req.Email)
	if err != nil {
		return ErrAuthorNotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := s.tokenService.Encode(author)
	if err != nil {
		return err
	}
	res.Token = token
	res.Valid = true

	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *proto.Token, res *proto.Token) error {
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.Author.Id == "" {
		return ErrInvalidAuthor
	}

	res.Valid = true

	return nil
}
