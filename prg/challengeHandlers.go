package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gurkslask/kaggaoc/sqlc/kaggaoc"
)

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	// Get which challenge
	q := r.URL.Query()
	c := q.Get("challenge")
	var err error
	templateURL := fmt.Sprintf("templates/challenge%v.html", c)
	fmt.Println(templateURL)
	t, err := template.ParseFS(templateFS, templateURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//http.Error(w, fmt.Sprintf("Challenge %v does not exist", c), http.StatusInternalServerError)
		return
	}
	// SessionData
	session, _ := Store.Get(r, "your-session-name")
	auth, _ := session.Values["authenticated"].(bool)
	username, _ := session.Values["user_name"].(string)

	data := struct {
		Authenticated bool
		Username      string
		Challenge     string
	}{auth, username, c}
	// Execute page
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func inputChallengeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	// Get which challenge
	q := r.URL.Query()
	c := q.Get("challenge")
	ci, err := strconv.Atoi(c)
	if err != nil {
		http.Error(w, "Illegal challenge", http.StatusInternalServerError)
	}
	//t, err := template.ParseFS(templateFS, "templates/challenge1_input.html")
	templateURL := fmt.Sprintf("templates/challenge1_input.html")
	t, err := template.ParseFS(templateFS, templateURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := Store.Get(r, "your-session-name")
	auth, _ := session.Values["authenticated"].(bool)
	username, _ := session.Values["user_name"].(string)

	p, err := gProblems.GetProblem(ci)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.SetSeed(session.Values["seed"].(int64))
	p.GenerateInputAndAnswer()

	data := struct {
		Authenticated bool
		Username      string
		InputData     string
		Challenge     string
	}{auth, username, p.GetInput(), c}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func answerChallengeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var err error
		// Get which challenge
		q := r.URL.Query()
		c := q.Get("challenge")
		ci, err := strconv.Atoi(c)
		if err != nil {
			http.Error(w, "Illegal challenge", http.StatusInternalServerError)
		}
		fmt.Sprintf("Check answer\n")
		r.ParseForm()
		answer := r.FormValue("answer")
		t, err := template.ParseFS(templateFS, "templates/challenge1_check.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := Store.Get(r, "your-session-name")
		auth, _ := session.Values["authenticated"].(bool)
		username, _ := session.Values["user_name"].(string)

		p, err := gProblems.GetProblem(ci)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.SetSeed(session.Values["seed"].(int64))
		p.GenerateInputAndAnswer()
		fmt.Sprintf("Check answer, your answer is: %v, true answer: %v\n", answer, p.GetAnswer())
		if answer == p.GetAnswer() {
			userID, err := queries.GetUserId(ctx, username)
			puserID := int32ToInt4(userID)
			_, err = queries.CreateChallenge(ctx, kaggaoc.CreateChallengeParams{
				Challenge: int32(ci),
				UserID:    puserID})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

		data := struct {
			Authenticated bool
			Username      string
			Trueanswer    string
			Answer        string
			Challenge     string
		}{auth, username, p.GetAnswer(), answer, c}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func challengesHandler(w http.ResponseWriter, r *http.Request) {
	// Get which challenge
	q := r.URL.Query()
	c := q.Get("challenge")
	var err error
	templateURL := fmt.Sprintf("templates/challenges.html")
	fmt.Println(templateURL)
	t, err := template.ParseFS(templateFS, templateURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//http.Error(w, fmt.Sprintf("Challenge %v does not exist", c), http.StatusInternalServerError)
		return
	}
	// SessionData
	session, _ := Store.Get(r, "your-session-name")
	auth, _ := session.Values["authenticated"].(bool)
	username, _ := session.Values["user_name"].(string)
	// Make challenge map
	type s struct {
		Num       int
		Desc      string
		Completed bool
	}
	problemSlice := strings.Split(gProblems.GetProblems(), "\n")
	//_, completedChallenges := getCompletedChallenges(db, username)
	userID, err := queries.GetUserId(ctx, username)
	puserID := int32ToInt4(userID)
	completedChallenges, err := queries.GetChallengeCompleted(ctx, puserID)
	ss := []s{}
	// Remove the last one
	for n, prob := range problemSlice[:len(problemSlice)-1] {
		st := s{n + 1, prob, ContainsInt32(completedChallenges, int32(n))}
		ss = append(ss, st)
	}
	data := struct {
		Authenticated bool
		Username      string
		Challenge     string
		Challenges    []s
	}{auth, username, c, ss}
	// Execute page
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
