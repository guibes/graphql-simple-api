package database

import (
	"context"
	"log"
	"time"

	"github.com/guibes/graphql-simple-api/graph/model"
	"github.com/guibes/graphql-simple-api/graph/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27100"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db *DB) Save(input *model.NewUser) *model.User {
	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}
	input.Password = passwordHash

	input.CreatedAt = int(time.Now().Unix())
	collection := db.client.Database("users").Collection("login_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.User{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
}

func (db *DB) FindByID(ID string) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("users").Collection("login_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"id": ObjectID})

	user := *&model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) FindUserByEmail(Email string) (*model.User, error) {
	collection := db.client.Database("users").Collection("login_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"email": Email})

	user := *&model.User{}
	res.Decode(&user)
	return &user, nil
}

func (db *DB) All() []*model.User {
	collection := db.client.Database("users").Collection("login_details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var users []*model.User

	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
