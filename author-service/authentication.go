package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/tamarakaufler/publication-manager/author-service/proto"
)

var (

	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("b9d4d807c28ca2b63bc7468e350e4493")
)

// hashed and sent as the second segment of the JWT
type Claims struct {
	Author *proto.Author
	jwt.StandardClaims
}

type Authentication interface {
	Decode(token string) (*Claims, error)
	Encode(author *proto.Author) (string, error)
}

type TokenService struct {
}

func (ts TokenService) Decode(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token => return the custom claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (ts TokenService) Encode(author *proto.Author) (string, error) {

	expireToken := time.Now().Add(time.Hour * 48).Unix()

	claims := Claims{
		author,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "publication-manager.author",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signing the token
	return token.SignedString(key)
}
