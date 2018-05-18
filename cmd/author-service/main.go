package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	_ "github.com/micro/go-plugins/registry/kubernetes"

	pb "github.com/tamarakaufler/publication-manager/author-service/proto"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
	defaultFilename = "author.json"
)

func parseFile(file string) (*pb.Author, error) {
	var author *pb.Author
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &author)
	return author, err
}

func main() {
	cmd.Init()

	client := pb.NewAuthorServiceClient("publication.management.author", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	author, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateAuthor(context.Background(), author)
	if err != nil {
		log.Fatalf("Could not create author: %v", err)
	}
	log.Printf("Author created: %t", r.Created)

	all, err := client.GetAuthors(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range all.Authors {
		log.Println(v)
	}
}
