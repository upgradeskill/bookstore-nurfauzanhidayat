package domain

import "github.com/golang-jwt/jwt"

type Users struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
	Token    string `json:"token,omitempty"`
}

type UserClaims struct {
	Userid  string `json:"userid"`
	Profile string `json:"profile"`
	Admin   bool   `json:"admin"`
	jwt.StandardClaims
}
