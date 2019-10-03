package main

import (
	pb "go-microservices/user-service/proto/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50053"
)

func main() {
	// Create postgresql db connection
	db, err := CreateConnection()

	if err != nil {
		log.Fatalf("Error occured while connecting to PostgresDB: %v \n", err)
	}

	// create user schema in postgres
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenservice := &TokenService{repo}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error occured, %v \n", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &service{repo, tokenservice})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
