package main

import (
	"errors"
	"fmt"
	pb "go-microservices/user-service/proto/user"
	"log"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := srv.repo.Get(req.Id)

	if err != nil {
		log.Fatalf("Error while fetching user: %v \n", err)
		return nil, err
	}

	return &pb.Response{User: user}, nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	users, err := srv.repo.GetAll()

	if err != nil {
		log.Fatalf("Error while fetching all users: %v \n", err)
		return nil, err
	}

	return &pb.Response{Users: users}, nil
}

func (srv *service) Auth(context context.Context, req *pb.User) (*pb.Token, error) {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := srv.repo.GetByEmailAndPassword(req)
	log.Println(user)
	if err != nil {
		return nil, err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (srv *service) Create(context context.Context, req *pb.User) (*pb.Response, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return nil, errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.tokenService.Encode(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: req, Token: &pb.Token{Token: token}}, nil
}

func (srv *service) ValidateToken(context context.Context, req *pb.Token) (*pb.Token, error) {
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	if claims.User.Id == "" {
		return nil, errors.New("invalid user")
	}
	return &pb.Token{Valid: true}, nil
}
