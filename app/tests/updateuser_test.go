package tests

import (
	"bytes"
	"context"
	"fmt"
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

func TestUpdateUser(t *testing.T) {

	usersCollection := db.Initialize().Collection("Users")
	userId := primitive.NewObjectID()
	newUser := user.User{
		ID:          userId,
		Name:        "Siddhesh",
		Dob:         "10 july",
		Address:     "Mangaon",
		Description: "Great user",
		CreatedAt:   time.Now().String(),
	}

	_, err := usersCollection.InsertOne(context.TODO(), newUser)

	util.HandleTestError(err, t)
	t.Cleanup(func() {
		usersCollection.DeleteOne(context.TODO(), bson.M{"_id": userId})
	})

	url := "/user/" + userId.Hex()

	router := app.CreateRouter()

	type UpdateOperation int
	const (
		Name UpdateOperation = iota
		Dob
		Address
		Description
	)
	tests := []struct {
		name            string
		updateField     string
		updateValue     string
		updateUperation UpdateOperation
	}{
		{name: "Updated name", updateField: "Name", updateValue: "Sahil", updateUperation: Name},
		{name: "Updated date of birth", updateField: "Dob", updateValue: "20 June", updateUperation: Dob},
		{name: "Updated address", updateField: "Address", updateValue: "Vinhere Mahad", updateUperation: Address},
		{name: "Updated description", updateField: "Description", updateValue: "Great boy and topper in class", updateUperation: Description},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody := fmt.Sprintf("{\"%v\":\"%v\"}", tt.updateField, tt.updateValue)
			util.HandleTestError(err, t)
			req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(jsonBody)))
			rr := httptest.NewRecorder()
			util.HandleTestError(err, t)
			router.ServeHTTP(rr, req)
			if rr.Result().StatusCode != http.StatusOK {
				t.Fatal(rr.Result().Status)
			}
			query := bson.M{"_id": userId}
			var updatedUser user.User
			findOneRes := usersCollection.FindOne(context.TODO(), query)
			if err := findOneRes.Err(); err != nil {
				util.HandleTestError(err, t)
			}
			findOneRes.Decode(&updatedUser)
			var valToCompare interface{}
			switch tt.updateUperation {
			case Name:
				valToCompare = updatedUser.Name
			case Dob:
				valToCompare = updatedUser.Dob
			case Address:
				valToCompare = updatedUser.Address
			case Description:
				valToCompare = updatedUser.Description
			}

			if tt.updateValue != valToCompare {
				t.Fatal(tt.updateField + " is different")
			}

		})
	}

	usersCollection.DeleteOne(context.TODO(), bson.M{"_id": userId})
}
