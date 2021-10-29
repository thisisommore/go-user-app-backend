package user

type BareUser struct {
	Name        string `json:"name" validate:"required"`
	Dob         string `json:"dob" validate:"datetime=05-Jan-2006"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type AddUserRequest BareUser

type AddUserResponse struct {
	UserId string `json:"userId"`
}

type UpdateUserRequest struct {
	Name        string `bson:"name,omitempty"`
	Dob         string `bson:"dob,omitempty"`
	Address     string `bson:"address,omitempty"`
	Description string `bson:"description,omitempty"`
}

const Layout = "05-Jan-2006"
