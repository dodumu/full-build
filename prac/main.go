package main

import (
	"net/http"
	"user/helper"
)

func main() {
	http.HandleFunc("/users/", helper.GetUserByID)
	http.ListenAndServe(":8081", nil)
}
