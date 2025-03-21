package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	"fmt"

	"github.com/gurkslask/kaggaoc/sqlc/kaggaoc"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("POSTTT")

		// Hämta data från formuläret
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}

		// Validera data (t.ex. kontrollera längd, format, etc.)

		// Skapa användaren i databasen (använd createUser-funktionen från tidigare)
		_, err = queries.CreateUser(ctx, kaggaoc.CreateUserParams{
			Username:     username,
			PasswordHash: string(hashedPassword[:]),
			Email:        email,
			Seed:         strconv.FormatInt(rand.Int63(), 10)})
		// err := createUser(db, username, password, email)
		fmt.Println(err)
		if err != nil {
			// Hantera fel, t.ex. visa ett felmeddelande
			fmt.Printf("FEL1: %v", err)
			http.Error(w, "Kunde inte skapa användare", http.StatusInternalServerError)
			return
		}
	}
	fmt.Println("HEHE")
	t, err := template.ParseFS(templateFS, "templates/register.html")
	if err != nil {
		fmt.Printf("FEL2: %v", err)
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
		fmt.Printf("FEL3: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("POSTTT")

		// Hämta data från formuläret
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		var user kaggaoc.User
		var err error

		userId, err := queries.GetUserId(ctx, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = queries.GetUser(ctx, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		username = user.Username
		hashedPassword := user.PasswordHash
		seed, _ := strToInt64(user.Seed)

		// Jämför det inmatade lösenordet med det hashade lösenordet
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			// Felaktigt lösenord
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		session, _ := Store.Get(r, "your-session-name")
		session.Values["user_id"] = userId
		session.Values["user_name"] = username
		session.Values["authenticated"] = true
		session.Values["seed"] = seed
		session.Save(r, w)
	}

	t, err := template.ParseFS(templateFS, "templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "your-session-name")
	_, ok := session.Values["user_id"].(int)
	if !ok {
		// Användaren är inte inloggad
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Användaren är inloggad, gör något med userId
}

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//session, _ := store.Get(r, "your-session-name")
		//session.Values["authenticated"]
		// ... gör något med sessionen ...
		next.ServeHTTP(w, r)
	})
}
