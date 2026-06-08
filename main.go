package main

import (
	"net/http"
	"user/helper"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/users/", helper.GetUserByID)
	http.HandleFunc("/", helper.Handler)
	http.HandleFunc("/health", helper.HealthHandler)
	http.HandleFunc("/create-user", helper.CreateUser)
	http.HandleFunc("/edit/", helper.EditUser)
	http.HandleFunc("/update/", helper.UpdateUser)
	http.HandleFunc("/delete/", helper.DeleteUser)
	http.ListenAndServe(":8081", nil)
}
