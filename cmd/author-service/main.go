package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/cmd"
	microclient "github.com/micro/go-micro/client"

	pb "github.com/tamarakaufler/publication-manager/author-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	// Create new greeter client
	client := pb.NewAuthorServiceClient("publication.management.author", microclient.DefaultClient)


	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Error connecting: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewAuthorServiceClient(conn)

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
