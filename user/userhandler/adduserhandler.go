package userhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request user.AddUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	util.HandleError(err)

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// TODO: Add log, plan logging, check internshala
	// log.Println(err)
	if err := validate.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		for _, err := range err.(validator.ValidationErrors) {
			var errMsg string
			if err.Tag() == "required" {
				errMsg = fmt.Sprintf("%v is required\n", err.Field())
			} else {
				errMsg = fmt.Sprintf("%v is invalid, needed %v\n", err.Field(), err.Tag())
			}
			w.Write([]byte(errMsg))
		}
		return
	}

	//TODO: docs to specify datetime format
	util.HandleError(err)
	newUser := user.User{
		Name:        request.Name,
		Dob:         request.Dob,
		Address:     request.Address,
		Description: request.Description,
	}
	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	//TODO: Better error strategy
	res, err := db.Db.Collection("Users").InsertOne(context.TODO(), newUser)
	util.HandleError(err)
	jsonResponse, _ := json.Marshal(user.AddUserResponse{UserId: res.InsertedID.(primitive.ObjectID).Hex()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
