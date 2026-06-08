package helper

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var dataFile = "user.json"

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := LoadUsers(dataFile)
	if err != nil {
		users = []User{}
	}

	// err = tmpl.Execute(w, users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	data := PageData{
		Users: users,
		Count: len(users),
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := LoadUsers(dataFile)
	if err != nil {
		users = []User{}
	}

	renderErrorPage := func(msg string) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := PageData{
			Users: users,
			Count: len(users),
			Error: msg,
		}

		tmpl.Execute(w, data)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	age := r.FormValue("age")

	userAge, err := strconv.Atoi(age)
	if err != nil {
		renderErrorPage("Age must be a valid number")
		return
	}

	if name == "" {
		renderErrorPage("Name is required")
		return
	}

	if email == "" {
		renderErrorPage("Email is required")
		return
	}

	if userAge <= 0 {
		renderErrorPage("Age must be greater than 0")
		return
	}

	newUser := User{
		ID:    GetNextID(users),
		Name:  name,
		Email: email,
		Age:   userAge,
	}

	err = CheckUserEmail(users, newUser)
	if err != nil {
		renderErrorPage(err.Error())
		return
	}

	users = AddUser(users, newUser)

	err = SaveUsers(dataFile, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/edit/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, err := LoadUsers(dataFile)
	if err != nil {
		users = []User{}
	}
	target := FindUserByID(users, id)
	if target == nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := LoadUsers(dataFile)
	if err != nil {
		// users = []User{}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	path := r.URL.Path
	path = strings.TrimPrefix(path, "/update/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	target := FindUserByID(users, id)
	if target == nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	newName := r.FormValue("name")
	newAge := r.FormValue("age")
	newAgeInt, err := strconv.Atoi(newAge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newEmail := r.FormValue("email")

	*target = User{
		ID:    id,
		Name:  newName,
		Age:   newAgeInt,
		Email: newEmail,
	}
	err = SaveUsers(dataFile, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := LoadUsers(dataFile)
	if err != nil {
		// users = []User{}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	path := r.URL.Path
	path = strings.TrimPrefix(path, "/delete/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, ok := RemoveUserByID(users, id)
	if !ok {
		http.Error(w, "Unable to delete user", http.StatusNotModified)
		return
	}

	err = SaveUsers(dataFile, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
