package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Dob         string             `bson:"dob"`
	Address     string             `bson:"address"`
	Description string             `bson:"description"`
	CreatedAt   string             `bson:"createdAt"`
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
