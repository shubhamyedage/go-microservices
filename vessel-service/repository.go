package main

import (
	"context"
	"errors"
	"log"

	pb "go-microservices/vessel-service/proto/vessel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository -
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

//VesselRepository -
type VesselRepository struct {
	collection *mongo.Collection
}

//FindAvailable -
func (vesselRepository *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	log.Println("In FindAvailable")
	cur, _ := vesselRepository.collection.Find(context.Background(), bson.D{{}}, nil)

	var vessels []*pb.Vessel
	for cur.Next(context.Background()) {
		var vessel *pb.Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		vessels = append(vessels, vessel)
	}

	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

//Create -
func (vesselRepository *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := vesselRepository.collection.InsertOne(context.Background(), vessel)
	return err
}
