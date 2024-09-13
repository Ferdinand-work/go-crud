package services

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/Ferdinand-work/go-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	req, _ := http.NewRequest("GET", os.Getenv("EXT_API"), nil)

	req.Header.Add("x-rapidapi-key", "Sign Up for Key")
	req.Header.Add("x-rapidapi-host", "easy-fast-temp-mail.p.rapidapi.com")

	resEmail, _ := http.DefaultClient.Do(req)
	email, _ := io.ReadAll(resEmail.Body)
	defer resEmail.Body.Close()
	user.Email = string(email)

	res, err := u.userCollection.InsertOne(u.ctx, user)
	return res, err
}
func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}
func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	defer cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}

	return users, nil
}
func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}, bson.E{Key: "email", Value: user.Email}}}}
	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) GetByAge(age int64) ([]*models.User, error) {
	var users []*models.User
	query := bson.D{bson.E{Key: "user_age", Value: age}}
	cursor, err := u.userCollection.Find(u.ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	defer cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}

	return users, nil
}

func (u *UserServiceImpl) AddFriends(usernames *[]string, name string) (*[]string, error) {

	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "friends", Value: usernames}}}}
	result, err := u.userCollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return nil, errors.New("cannot update")
	}
	if result.MatchedCount < 1 {
		return nil, errors.New("no matched document found for update")
	}
	return usernames, nil
}

func (u *UserServiceImpl) GetFriends(name string) (*[]models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return nil, err
	}
	friendNames := user.Friends
	filter := bson.M{"user_name": bson.M{"$in": &friendNames}}
	cursor, err := u.userCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	var pointer = &users
	if err = cursor.All(context.Background(), pointer); err != nil {
		return nil, err
	}
	return pointer, nil
}
