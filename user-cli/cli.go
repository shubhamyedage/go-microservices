package main

import (
	"context"
	"log"

	pb "go-microservices/user-service/proto/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50053"
)

func main() {
	// Setup grpc connection to user-service
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while connecting. %v \n", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	log.Println("Running user service on address: %v.", address)

	// Create user
	res, err := client.Create(context.Background(), &pb.User{
		Name:     "Synerzip",
		Email:    "synerzip@synerzip.com",
		Password: "synerzip",
		Company:  "synerzip",
	})

	if err != nil {
		log.Fatalf("Faild to create user %v \n", err)
	}

	log.Println("User created: ", res.User.Name)

	// List all users
	res, err = client.GetAll(context.Background(), &pb.Request{})

	if err != nil {
		log.Fatalf("Faild to create user %v \n", err)
	}

	log.Println("Users --->")
	for _, v := range res.Users {
		log.Println(v.Name)
	}

}
