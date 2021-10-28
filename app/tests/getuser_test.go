package tests

import (
	"context"
	"encoding/json"
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
)

func TestGetUser(t *testing.T) {

	usersCollection := db.Initialize().Collection("Users")
	userId := primitive.NewObjectID()
	newUser := user.User{
		ID:          userId,
		Name:        "Tom",
		Dob:         "10 july",
		Address:     "Mangaon",
		Description: "Great user",
		CreatedAt:   time.Now().String(),
	}
	_, err := usersCollection.InsertOne(context.TODO(), newUser)
	util.HandleTestError(err, t)

	url := "/user/" + userId.Hex()
	rr := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", url, nil)

	router := app.CreateRouter()
	router.ServeHTTP(rr, request)
	var userGot user.User
	json.Unmarshal(rr.Body.Bytes(), &userGot)
	if !user.AreUsersEqual(newUser, userGot) {
		t.Fatal("User data is not as expected")
	}
	usersCollection.DeleteOne(context.TODO(), bson.M{"_id": userId})
}
