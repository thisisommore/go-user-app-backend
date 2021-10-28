package userhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateUserRequest user.User

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	usersCollection := db.Db.Collection("Users")

	userId := mux.Vars(r)["id"]
	var body interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)
	_, err := bson.Marshal(body)
	updateFilter := bson.M{"$set": body}
	util.HandleError(err)
	objectId, _ := primitive.ObjectIDFromHex(userId)
	updateRes, err := usersCollection.UpdateByID(context.TODO(), objectId, updateFilter)
	util.HandleError(err)
	if updateRes.MatchedCount == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
