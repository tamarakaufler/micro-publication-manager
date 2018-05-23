package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	proto "github.com/tamarakaufler/publication-manager/book-service/proto"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "book.json"
)

func parseFile(file string) (*proto.Book, error) {
	var book *proto.Book
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &book)
	return book, err
}

func main() {
	cmd.Init()

	client := proto.NewBookServiceClient("publication.manager.book", microclient.DefaultClient)

	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}

	file = os.Args[1]
	token = os.Args[2]

	book, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// Test 1
	r, err := client.PublishBook(ctx, book)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Registered: %t", r.Registered)

	// Test 2
	all, err := client.GetPublishedBooks(ctx, &proto.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list books: %v", err)
	}
	for _, v := range all.Books {
		log.Println(v)
	}
}
