package main

import (
	"net/http"
	"user/helper"
)

func main() {
	http.HandleFunc("/users/", helper.GetUserByID)
	http.HandleFunc("/", helper.Handler)
	http.HandleFunc("/health", helper.HealthHandler)
	http.HandleFunc("/create-user", helper.CreateUser)
	http.HandleFunc("/edit/", helper.EditUser)
	http.ListenAndServe(":8081", nil)
}
