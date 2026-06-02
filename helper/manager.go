package helper

import (
	"encoding/json"
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

func RemoveUserByID(users []User, id int) ([]User, bool) {
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return append(users[:i], users[i+1:]...), true
		}
	}
	return users, false
}

func CheckUserEmail(users []User, newUser User) error {

	for i := 0; i < len(users); i++ {
		if users[i].Email == newUser.Email {
			return errors.New("email already exists")
		}
	}
	return nil
}

func FindUserByID(users []User, id int) *User {
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return &users[i]
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
	dataFile := "user.json"
	// w.Header().Set("Content-Type", "application/json")
	// err = json.NewEncoder(w).Encode(target)
	// if err != nil {
	// 	http.Error(w, "user not found", http.StatusInternalServerError)
	// 	return
	// }

	switch r.Method {
	case http.MethodDelete:
		user, err := LoadUsers(dataFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, found := RemoveUserByID(user, id)
		if !found {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		err = SaveUsers(dataFile, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodPut:
		user, err := LoadUsers(dataFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		target := FindUserByID(user, id)

		if target == nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		var updatedUser User
		err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		target.Name = updatedUser.Name
		target.Email = updatedUser.Email
		target.Age = updatedUser.Age

		err = SaveUsers(dataFile, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(target)
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
