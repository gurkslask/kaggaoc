package main

import (
	"database/sql"
	"math/rand"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// User representerar en användare i databasen
type User struct {
	ID           int
	Username     string
	PasswordHash string
	Email        string
}

func createUser(db *sql.DB, username, password, email string) error {
	// Hash lösenordet
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Förbered SQL-satser
	_, err = db.Exec("INSERT INTO users (username, password_hash, email, seed) VALUES ($1, $2, $3, $4)", username, hashedPassword, email, rand.Int63())
	if err != nil {
		return err
	}

	return nil
}
