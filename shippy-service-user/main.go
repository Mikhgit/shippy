package main

import (
	"log"

	pb "github.com/Mikhgit/shippy/shippy-service-user/proto/auth"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/broker/nats/v2"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := NewConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()

	// Run schema query on start-up, as we're using "create if not exists"
	// this will only be ran once. In order to create updates, you'll need to
	// use a migrations library
	db.MustExec(schema)

	repo := NewPostgresRepository(db)

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("go.micro.api.auth"),
		micro.Version("latest"),
		//	micro.Broker(nats.NewBroker()),
	)

	// Init will parse the command line flags.
	service.Init()

	publisher := micro.NewEvent("user.created", service.Client())

	// Register handler
	if err := pb.RegisterAuthHandler(service.Server(), &handler{
		repo, tokenService, publisher,
	}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
