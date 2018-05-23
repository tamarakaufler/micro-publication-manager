package main

import (
	"log"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	k8s "github.com/micro/kubernetes/go/micro"
	proto "github.com/tamarakaufler/publication-manager/author-service/proto"
)

func main() {
	log.Println("Starting author-service ...")

	conn, err := DBConnection()
	defer conn.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		log.Println("Connected to database ...")
	}

	// Automatically migrates author struct
	// into database columns/types etc
	// Will migrate changes each time
	// the service is restarted.
	conn.AutoMigrate(&proto.Author{})
	db := &Store{conn}
	tokenService := TokenService{}

	microSrv := k8s.NewService(
		micro.Name("publication.management.author"),
		micro.Version("latest"),
	)
	microSrv.Init()

	proto.RegisterAuthorServiceHandler(microSrv.Server(), &service{db, tokenService})

	//reflection.Register(s)
	if err := microSrv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
