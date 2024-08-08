package data

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/initializers"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User = model.User

var userCollection *mongo.Collection
var jwtSecret = []byte("your_jwt_secret")

func getUserCollection() *mongo.Collection {
	if userCollection == nil {
		userCollection = initializers.Client.Database("test").Collection("users")
	}
	return userCollection
}

// CreateUser inserts a new user into the database.
func CreateUser(user User) (User, error) {
	userCollection := getUserCollection()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPassword)

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

func LoginUser(user User) (string, error) {
	userCollection := getUserCollection()
	// fetch the user from the database
	fillter := bson.D{{"email", user.Email}}
	var newUser User
	err := userCollection.FindOne(context.TODO(), fillter).Decode(&newUser)

	if err != nil {
		return "", err
	}

	//compare the hased password with the user given password
	if bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)) != nil {
		return "", err
	}

	// generet the jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    newUser.ID,
		"email": user.Email,
		"role":  newUser.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	// return the token
	return jwtToken, nil
}

func PromoteUser(userId string) (User, error) {
	// get the user collection
	userCollection := getUserCollection()
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
