package helper

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {

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
	RenderTemplate(w, "index.html", data)
}
