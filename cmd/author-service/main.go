package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	proto "github.com/tamarakaufler/publication-manager/author-service/proto"
	"golang.org/x/net/context"
)

func main() {
	microSrv := micro.NewService(
		micro.Name("publication.manager.author-cli"),
		micro.Version("latest"),
	)
	microSrv.Init()

	client := proto.NewAuthorServiceClient("publication.manager.author", microclient.DefaultClient)

	firstName := "Lucien"
	lastName := "Kaufler"
	email := "lucien@gmail.com"
	password := "topsecreteverybodyknows"

	r, err := client.CreateAuthor(context.TODO(), &proto.Author{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.Author.Id)

	all, err := client.GetAll(context.Background(), &proto.GetAllRequest{})
	if err != nil {
		log.Fatalf("Could not list authors: %v", err)
	}
	for _, v := range all.Authors {
		log.Println(v)
	}

	authResponse, err := client.Authenticate(context.TODO(), &proto.Author{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate author: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)
	os.Exit(0)
}
