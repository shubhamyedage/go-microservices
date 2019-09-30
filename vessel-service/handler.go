package main

import (
	"context"

	pb "go-microservices/vessel-service/proto/vessel"
)

type handler struct {
	repo Repository
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification) (*pb.Response, error) {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Vessel: vessel}, nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel) (*pb.Response, error) {
	err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true}, nil
}
