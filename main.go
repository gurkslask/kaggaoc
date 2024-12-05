package main

import (
        "database/sql"
        "embed"
        "html/template"
        "net/http"

        _ "github.com/lib/pq"
)

//go:embed templates/*
var templateFS embed.FS

// ... (definiera strukturer för användare, utmaningar, etc.)

func main() {
        db, err := sql.Open("postgres", "host=db user=postgres password=mysecretpassword dbname=advent_of_code")
        if err != nil {
                panic(err)
        }
        defer db.Close()

        // ... (skapa tabeller i databasen)

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                t, err := template.ParseFS(templateFS, "templates/index.html")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                // ... (hämta data från databasen och skicka till template)
                data := ""

                err = t.Execute(w, data)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                }
        })

        http.ListenAndServe(":8080", nil)
}
