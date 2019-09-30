package main

import (
	pb "go-microservices/consignment-service/proto/consignment"
	vesselProto "go-microservices/vessel-service/proto/vessel"
	"log"

	"golang.org/x/net/context"
)

//Handler -
type handler struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	log.Println("In CreateConsignment")
	vesselResp, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResp.Vessel.Name)
	if err != nil {
		return nil, err
	}
	req.VesselId = vesselResp.Vessel.Id

	// Save our consignment
	er := s.repo.Create(req)
	if er != nil {
		return nil, er
	}

	return &pb.Response{Created: true}, nil
}

// GetConsignments - list all consignments.
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	log.Println("In GetConsignments()")
	log.Println("Get all consignments....")
	consignments, _ := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}
