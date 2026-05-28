package main

import (
	"fmt"
	"user/user"
)

func main() {
	person := user.User{
		Name:  "Abigail",
		Age:   29,
		Email: "Aq2@gmail.com",
	}

	users, err := user.LoadUsers("some.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.FindUser(users, "davidrobertq2@gmail.com"))

	users = user.AddUser(users, person)

	user.SaveUsers("some.json", users)

	fmt.Println(user.RemoveUser(person, users))
}
