package main

import (
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (t User) Display() {
	fmt.Printf("Name: %v\nAge: %v\nEmail: %v\n", t.Name, t.Age, t.Email)
}

func (t *User) Birthday() {
	t.Age += 1
}

func (t *User) UpdateEmail(newEmail string) {
	t.Email = newEmail

}

func (t *User) Rename(newName string) {
	t.Name = newName
}

func main() {
	person := User{
		Name:  "David",
		Age:   29,
		Email: "davidrobertq2@gmail.com",
	}

	person.Display()
	person.Birthday()
	person.Display()
	person.UpdateEmail("brucenic12@gmail.com")
	person.Display()
	person.Rename("Lucky")
	person.Display()
}
