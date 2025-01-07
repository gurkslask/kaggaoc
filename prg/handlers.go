package main

import (
	"fmt"
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

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	t, err := template.ParseFS(templateFS, "templates/challenge1.html")
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
func inputChallengeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	t, err := template.ParseFS(templateFS, "templates/challenge1_input.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := Store.Get(r, "your-session-name")
	auth, _ := session.Values["authenticated"].(bool)
	username, _ := session.Values["user_name"].(string)

	var p Problem
	p.Seed = session.Values["seed"].(int64)
	p.GenerateInputAndAnswer()

	data := struct {
		Authenticated bool
		Username      string
		InputData     string
	}{auth, username, p.Input}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func answerChallengeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Sprintf("Check answer\n")
		r.ParseForm()
		answer := r.FormValue("answer")
		var err error
		t, err := template.ParseFS(templateFS, "templates/challenge1_check.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := Store.Get(r, "your-session-name")
		auth, _ := session.Values["authenticated"].(bool)
		username, _ := session.Values["user_name"].(string)

		var p Problem
		p.Seed = session.Values["seed"].(int64)
		p.GenerateInputAndAnswer()
		fmt.Sprintf("Check answer, your answer is: %v, true answer: %v\n", answer, p.Answer)

		data := struct {
			Authenticated bool
			Username      string
			Trueanswer    string
			Answer        string
		}{auth, username, p.Answer, answer}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
