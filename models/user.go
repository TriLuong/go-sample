package models

type (
	User struct {
		Role      string `json:"role"`
		Email     string `json:"email"`
		Passoword string `json:"password"`
		Phone     string `json:"phone"`
		Token     string `json:"token"`
	}
)
