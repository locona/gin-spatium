package models

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Users []User

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MockUsers() Users {
	var users Users
	for i := 0; i < 10; i++ {
		name := fmt.Sprint("name_%d", i)
		user := User{
			ID:   i,
			Name: name,
		}
		pp.Println(user)
		users = append(users, user)
	}
	return users
}
