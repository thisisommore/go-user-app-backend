package userhandler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	usersCollection := db.Db.Collection("Users")

	userId := mux.Vars(r)["id"]
	if userId == "" {
		w.Write([]byte("id is required"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id is not valid"))
		return
	}
	delOneRes, err := usersCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	util.HandleError(err)
	if delOneRes.DeletedCount == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
