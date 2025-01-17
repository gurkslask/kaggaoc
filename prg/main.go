package main

import (
	"database/sql"
	"embed"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

//go:embed templates/*
var templateFS embed.FS

// Gorilla sessions
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var db *sql.DB

var username string

var gProblems ProblemStruct

func main() {
	gProblems = CreateProblemStruct()
	fs := http.FileServer(http.Dir("assets/"))
	username = ""
	var err error
	db, err = sql.Open("postgres", "host=db user=postgres password=mysecretpassword dbname=advent_of_code sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerPageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/challenge", challengeHandler)
	http.HandleFunc("/challenge_input", inputChallengeHandler)
	http.HandleFunc("/challenge_check", answerChallengeHandler)
	http.HandleFunc("/challenges", challengesHandler)

	http.ListenAndServe(":8080", nil)
}
