package main

import (
	pb "go-microservices/user-service/proto/user"
	"log"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		log.Fatalf("Faild to get users. %v\n", err)
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id

	if err := repo.db.Find(&user).Error; err != nil {
		log.Fatalf("Faild to get user by ID. %v\n", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		log.Fatalf("Faild to create user. %v\n", err)
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
