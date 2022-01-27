package _interface

import (
	"github.com/dgrijalva/jwt-go"
)

type Staff interface {
	Login() jwt.Token
	LogOut()
}
