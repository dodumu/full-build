package main

import (
	"net/http"
	"user/user"
)

func main() {
	http.HandleFunc("/", user.Handler)

	http.HandleFunc("/users", user.Hanlder1)

	http.ListenAndServe(":8081", nil)
}
