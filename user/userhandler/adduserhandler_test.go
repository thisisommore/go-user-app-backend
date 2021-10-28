package userhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUser(t *testing.T) {
	newUser := user.User{
		Name:        "Test Om",
		Dob:         "tuesday",
		Address:     "Goregaon, Mangaon, Raigad",
		Description: "I am a hardworking individual",
	}
	jsonBody, err := json.Marshal(newUser)
	//TODO
	util.HandleTestError(err, t)
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	util.HandleError(err)

	rr := httptest.NewRecorder()
	db.Initialize()
	AddUser(rr, req)
	if rr.Result().StatusCode != 200 {
		t.FailNow()
	}

	var res addUserResponse
	json.Unmarshal(rr.Body.Bytes(), &res)

	coll := db.Db.Collection("Users")
	obId, _ := primitive.ObjectIDFromHex(res.UserId)
	var addedUser user.User
	err = coll.FindOne(context.TODO(), bson.M{"_id": obId}).Decode(&addedUser)
	util.HandleTestError(err, t)

	if !user.AreUsersEqualIgnoringIdAndCreateAt(addedUser, newUser) {
		t.Fatalf("Added user doesn't match new user to be added")
	}

	delRes, err := coll.DeleteOne(context.TODO(), bson.M{"_id": obId})
	util.HandleTestError(err, t)
	if delRes.DeletedCount != 1 {
		t.Fatal("Failed to delete test user")
	}
}
