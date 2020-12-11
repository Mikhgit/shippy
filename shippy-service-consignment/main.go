// shippy-service-consignment/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"log"
	"os"
	"strings"

	pb "github.com/Mikhgit/shippy/shippy-service-consignment/proto/consignment"
	userService "github.com/Mikhgit/shippy/shippy-service-user/proto/auth"
	vesselProto "github.com/Mikhgit/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Authorization"]
		if len(strings.Fields(token)) != 2 {
			return errors.New("invalid token")
		}
		log.Println("Authenticating with token: ", strings.Fields(token)[1])

		// Auth here
		authClient := userService.NewAuthService("go.micro.api.auth", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: strings.Fields(token)[1],
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

func main() {
	// Set-up micro instance
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the service name given in your protobuf definition
		micro.Name("go.micro.api.consignment-service"),
		micro.Version("latest"),
		// Our auth middleware
		micro.WrapHandler(AuthWrapper),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("go.micro.api.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	err = pb.RegisterConsignmentServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
