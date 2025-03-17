package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templateFS, "templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := Store.Get(r, "your-session-name")

	auth, _ := session.Values["authenticated"].(bool)
	username, _ := session.Values["user_name"].(string)
	data := struct {
		Authenticated bool
		Username      string
	}{auth, username}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
