package main

import (
	"log"
	"net"
	"os"

	pb "go-microservices/vessel-service/proto/vessel"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":50052"
	defaultHost = "mongodb://localhost:27017"
)

func createDummyVesselData(repo Repository) {
	log.Println("In createDummyVesselData")
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessels")

	repository := &VesselRepository{vesselCollection}
	createDummyVesselData(repository)

	pb.RegisterVesselServiceServer(s, &handler{repository})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
