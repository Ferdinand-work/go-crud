package services

import (
	"github.com/Ferdinand-work/go-crud/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(*models.User) (*mongo.InsertOneResult, error)
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	GetByAge(int64) ([]*models.User, error)
	AddFriends(*[]string, string) (*[]string, error)
	GetFriends(string) (*[]models.User, error)
}
