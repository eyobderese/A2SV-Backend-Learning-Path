package repositories

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	infrastructure "github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Infrastructure"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtSecret = []byte("your_jwt_secret")

type User = domain.User

type userRepository struct {
	db             mongo.Database
	collectionName string
}

func NewUserRepository(db mongo.Database, collectionName string) domain.UserRepository {
	return &userRepository{db: db, collectionName: collectionName}
}

func (ur *userRepository) CreateUser(user User) (User, error) {
	userCollection := ur.db.Collection(ur.collectionName)
	hashedPassword, err := infrastructure.HashePassword(user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword

	// If the database is empty, the first user is an admin.
	count, err := userCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return User{}, err
	}
	if count == 0 {
		user.Role = "admin"
	}

	// if the it is not the first user, we check if the role if manully set else we set it to user

	if count > 0 && user.Role == "" {
		user.Role = "user"
	}

	_, err = userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return User{}, err
	}
	user.Password = ""
	return user, nil
}

func (ur *userRepository) LoginUser(user User) (string, error) {
	userCollection := ur.db.Collection(ur.collectionName)
	// fetch the user from the database
	fillter := bson.D{{"email", user.Email}}
	var newUser User
	err := userCollection.FindOne(context.TODO(), fillter).Decode(&newUser)

	if err != nil {
		return "", err
	}

	//compare the hased password with the user given password
	if !infrastructure.ComparePassword(newUser.Password, user.Password) {
		return "", err
	}

	// generet the jwt token

	token, err := infrastructure.TokenGeneretor(jwt.MapClaims{
		"id":    newUser.ID,
		"email": user.Email,
		"role":  newUser.Role,
	})

	// return the token
	return token, err
}

func (ur *userRepository) PromoteUser(userId string) (User, error) {
	// get the user collection
	userCollection := ur.db.Collection(ur.collectionName)
	// convert string Id to objectTypeId for the filter
	objectUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return User{}, err
	}

	//update the user role to admin
	filter := bson.D{{"_id", objectUserId}}
	update := bson.D{{"$set", bson.D{{"role", "admin"}}}}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return User{}, err
	}

	//return the updated user after removing the password
	var user User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return User{}, err
	}

	user.Password = ""
	return user, nil
}
