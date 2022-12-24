package main

import (
	"fmt"

	"github.com/aokuyama/go-study-layered/domain/model/user"
)

func main() {
	id, _ := user.NewUserID("abcdef")
	fmt.Println("Hello, User:", id)
}
