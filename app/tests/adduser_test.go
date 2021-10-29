package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thisisommore/go-user-app-backend/app"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUser(t *testing.T) {
	newUser := user.AddUserRequest{
		Name:        "Test Om",
		Dob:         "10-Jul-2002",
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
	router := app.CreateRouter()
	router.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		fmt.Println(rr.Body.String())
		t.FailNow()
	}

	var res user.AddUserResponse
	json.Unmarshal(rr.Body.Bytes(), &res)

	coll := db.Db.Collection("Users")
	obId, _ := primitive.ObjectIDFromHex(res.UserId)
	var addedUser user.User
	err = coll.FindOne(context.TODO(), bson.M{"_id": obId}).Decode(&addedUser)
	util.HandleTestError(err, t)

	const layout = "05-Jan-2006"
	newExpectedUser := user.User{
		ID:          obId,
		Name:        newUser.Name,
		Dob:         newUser.Dob,
		Address:     newUser.Address,
		Description: newUser.Description,
	}

	if !user.AreUsersEqualIgnoringIdAndCreateAt(addedUser, newExpectedUser) {
		t.Fatalf("Added user doesn't match new user to be added")
	}

	delRes, err := coll.DeleteOne(context.TODO(), bson.M{"_id": obId})
	util.HandleTestError(err, t)
	if delRes.DeletedCount != 1 {
		t.Fatal("Failed to delete test user")
	}
}
