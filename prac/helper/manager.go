package helper

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func AddUser(users []User, newUser User) []User {
	users = append(users, newUser)
	return users
}

func FindUserByEmail(users []User, email string) *User {
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return nil
}

func RemoveUserByID(users []User, id int) []User {
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}

func CheckUserEmail(users []User, newUser User) error {

	for i := 0; i < len(users); i++ {
		if users[i].Email == newUser.Email {
			return errors.New("email already exists")
		}
	}
	return nil
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	path = strings.TrimPrefix(path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// err = json.NewEncoder(w).Encode(target)
	// if err != nil {
	// 	http.Error(w, "user not found", http.StatusInternalServerError)
	// 	return
	// }

	switch r.Method {
	case http.MethodDelete:
		user, err := LoadUsers("user.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user = RemoveUserByID(user, id)
		err = SaveUsers("user.json", user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}

func ReturnUserByID(users []User, ID int) *User {
	for i := 0; i < len(users); i++ {
		if users[i].ID == ID {
			return (&users[i])
		}
	}
	return nil
}
