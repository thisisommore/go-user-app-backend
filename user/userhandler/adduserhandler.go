package userhandler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addUserRequest struct {
	Name        string
	Dob         string
	Address     string
	Description string
	//TODO dob,created check type
}

type AddUserResponse struct {
	UserId string `json:"userId"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request addUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	util.HandleError(err)

	// TODO: Add log, plan logging, check internshala
	// log.Println(err)
	newUser := user.User{
		Name:        request.Name,
		Dob:         request.Dob,
		Address:     request.Address,
		Description: request.Description,
	}
	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now().String()
	//TODO: Better error strategy
	res, err := db.Db.Collection("Users").InsertOne(context.TODO(), newUser)
	util.HandleError(err)
	jsonResponse, _ := json.Marshal(AddUserResponse{UserId: res.InsertedID.(primitive.ObjectID).Hex()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
