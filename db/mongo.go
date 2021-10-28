package db

import (
	"context"
	"os"

	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func Initialize() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	util.HandleError(err)
	Db = client.Database("userAppDatabase")
	return Db
}
