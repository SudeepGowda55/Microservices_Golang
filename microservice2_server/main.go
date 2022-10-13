package main

import (
	"fmt"
	"os"
	// "os"
	// "github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func main() {
	fmt.Println(mySigningKey)
}
