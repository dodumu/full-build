package helper

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	success := ""
	switch r.URL.Query().Get("success") {
	case "created":
		success = "User created successfully"
	case "delted":
		success = "User deleted"
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
		Success: success,
	}
	RenderTemplate(w, "index.html", data)
}
