package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/gurkslask/kaggaoc/sqlc/kaggaoc"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//go:embed templates/*
var templateFS embed.FS

// Gorilla sessions
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var username string

var l log.Logger

var gProblems ProblemStruct

var queries *kaggaoc.Queries

var ctx context.Context

func main() {
	var err error
	ctx = context.Background()

	f, err := os.OpenFile("/tmp/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	l.SetOutput(f)
	l.Println("Test")

	gProblems = CreateProblemStruct()
	fs := http.FileServer(http.Dir("assets/"))
	username = ""

	url := "postgres://postgres:mysecretpassword@db:5432/advent_of_code"
	db, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	queries = kaggaoc.New(db)

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
