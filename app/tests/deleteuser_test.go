package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/thisisommore/go-user-app-backend/app"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteUser(t *testing.T) {
	usersCollection := db.Initialize().Collection("Users")
	userId := primitive.NewObjectID()
	newUser := user.User{
		ID:          userId,
		Name:        "Doremon",
		Dob:         "10-July-2005",
		Address:     "Mangaon",
		Description: "Great user",
		CreatedAt:   time.Now(),
	}
	_, err := usersCollection.InsertOne(context.TODO(), newUser)
	util.HandleTestError(err, t)

	url := "/user/" + userId.Hex()
	router := app.CreateRouter()
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("DELETE", url, nil)
	util.HandleTestError(err, t)
	router.ServeHTTP(rr, request)
	if statusCode := rr.Result().StatusCode; statusCode != http.StatusOK {
		t.Fatal(statusCode)
	}

	findResErr := usersCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Err()
	if findResErr == nil {
		t.Fatal("User still exist")
	}
	if findResErr != mongo.ErrNoDocuments {
		t.Fatal(findResErr)
	}
	usersCollection.DeleteOne(context.TODO(), bson.M{"_id": userId})

}
