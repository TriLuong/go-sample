package models

type (
	User struct {
		Role      string `json:"role" bson:"role"`
		Email     string `json:"email" bson:"email"`
		Passoword string `json:"password" bson:"password"`
		Phone     string `json:"phone" bson:"phone"`
		Token     string `json:"token" bson:"token"`
	}
)
