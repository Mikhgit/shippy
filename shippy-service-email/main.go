// shippy-email-service
package main

import (
	"context"
	"log"

	pb "github.com/Mikhgit/shippy/shippy-service-user/proto/user"
	micro "github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/broker/nats/v2"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
		micro.Version("latest"),
	)

	srv.Init()

	err := micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err != nil {
		log.Println(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
