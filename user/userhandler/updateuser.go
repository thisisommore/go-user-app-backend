package userhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/thisisommore/go-user-app-backend/db"
	"github.com/thisisommore/go-user-app-backend/user"
	"github.com/thisisommore/go-user-app-backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	usersCollection := db.Db.Collection("Users")

	userId := mux.Vars(r)["id"]
	if userId == "" {
		w.Write([]byte("id is required"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	validate := validator.New()

	var body user.UpdateUserRequest

	json.NewDecoder(r.Body).Decode(&body)
	bsonMBytesBody, err := bson.Marshal(body)
	util.HandleError(err)
	var bsonMBody bson.M
	err = bson.Unmarshal(bsonMBytesBody, &bsonMBody)
	util.HandleError(err)
	validations := map[string]string{
		"name":        "required",
		"dob":         "datetime=05-Jan-2006",
		"address":     "required",
		"description": "required",
	}
	var errMsg string
	for k, v := range bsonMBody {
		validationTag := validations[k]
		if validationTag == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := validate.Var(v, validationTag)
		if err == nil {
			//There is no error, no need to write error message for current field
			continue
		}
		errs, _ := err.(validator.ValidationErrors)
		if errs[0].Tag() == "required" {
			errMsg += fmt.Sprintf("%v is required\n", k)
		} else {
			errMsg += fmt.Sprintf("%v is invalid, needed %v\n", k, errs[0].Tag())
		}
	}

	if errMsg != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMsg))
		return
	}
	updateFilter := bson.M{"$set": bsonMBody}
	objectId, _ := primitive.ObjectIDFromHex(userId)
	updateRes, err := usersCollection.UpdateByID(context.TODO(), objectId, updateFilter)
	util.HandleError(err)
	if updateRes.ModifiedCount == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
