package helper

import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type PageData struct {
	Users   []User
	Count   int
	Error   string
	Success string
}

func (t User) Display() {
	fmt.Println(t)
}

func (t *User) Birthday() {
	t.Age += 1
}

func (t *User) UpdateEmail(newEmail string) {
	t.Email = newEmail
}
