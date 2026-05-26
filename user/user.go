package user

import (
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
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
