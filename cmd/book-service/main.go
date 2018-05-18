package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	pb "github.com/tamarakaufler/publication-manager/book-service/proto"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "book.json"
)

func parseFile(file string) (*pb.Book, error) {
	var book *pb.Book
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &book)
	return book, err
}

func main() {
	cmd.Init()

	client := pb.NewBookServiceClient("publication.management.book", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	book, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.PublishBook(context.Background(), book)
	if err != nil {
		log.Fatalf("Could not publish book: %v", err)
		return
	}
	log.Printf("Book will be published by : %s", r.Book.GetPublisherId())

	all, err := client.GetPublishedBooks(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range all.Books {
		log.Println(v)
	}
}
