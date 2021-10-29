package user

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Dob         string             `bson:"dob"`
	Address     string             `bson:"address"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"createdAt"`
}

func AreUsersEqual(a User, b User) bool {
	if !AreUsersEqualIgnoringIdAndCreateAt(a, b) {
		fmt.Println("AreUsersEqualIgnoringIdAndCreateAt")
		return false
	}
	if a.ID != b.ID {
		fmt.Println("ID")

		return false
	}

	return true
}

func AreUsersEqualIgnoringIdAndCreateAt(a User, b User) bool {
	if a.Name != b.Name {
		return false
	}
	if a.Dob != b.Dob {
		return false
	}
	if a.Address != b.Address {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	return true
}
