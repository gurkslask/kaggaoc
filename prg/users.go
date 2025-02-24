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
func createChallengeDone(db *sql.DB, username string, challenge int) error {
	var userID int
	err := db.QueryRow("SELECT user_id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO completed (challenge, user_id) VALUES ($1, $2)", challenge, userID)
	if err != nil {
		return err
	}

	return nil
}
func getCompletedChallenges(db *sql.DB, username string) (error, []int) {
	var userID int
	err := db.QueryRow("SELECT user_id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return err, nil
	}
	rows, err := db.Query("SELECT challenge FROM completed WHERE user_id = $1", userID)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	var completed []int
	for rows.Next() {
		var comp int
		if err := rows.Scan(&comp); err != nil {
			return err, nil
		}
		completed = append(completed, comp)
	}
	return nil, completed

}
