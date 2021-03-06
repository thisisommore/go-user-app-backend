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

func GetUser(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["id"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is required"))
		return
	}
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id is not valid"))
		return
	}
	usersCollection := db.Db.Collection("Users")
	util.HandleError(err)

	filterQuery := bson.M{"_id": objectId}
	var userData user.User
	findRes := usersCollection.FindOne(context.TODO(), filterQuery)
	if err := findRes.Err(); err != nil {
		util.HandleError(err)
	}
	findRes.Decode(&userData)
	jsonResponse, err := json.Marshal(userData)
	util.HandleError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
