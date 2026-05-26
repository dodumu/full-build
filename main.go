package main

import (
	"encoding/json"
	"fmt"
	"os"
	"user/user"
)

func main() {
	person := user.User{
		Name:  "David",
		Age:   29,
		Email: "davidrobertq2@gmail.com",
	}

	val, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(string(val))

	err = os.WriteFile("user.json", val, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// person.Display()
	// person.Birthday()
	// person.Display()
	// person.UpdateEmail("brucenic12@gmail.com")
	// person.Display()
	// person.Rename("Lucky")
	// person.Display()
}
