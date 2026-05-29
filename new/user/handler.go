package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my backend")
}

func Hanlder1(w http.ResponseWriter, r *http.Request) {
	users, err := LoadUsers("user/some.json")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(users)
}
