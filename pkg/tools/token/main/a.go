package main

import (
	"github.com/XC-Zero/yinwan/pkg/tools/token"
	"log"
)

func main() {
	a := token.GenerateToken("123")
	log.Println(token.IsExpired(a))
}
