package main

import (
	"database/sql"

	"github.com/gurkslask/kaggaoc/sqlc/kaggaoc"
	_ "github.com/lib/pq"
)

// User representerar en användare i databasen
type User struct {
	ID           int
	Username     string
	PasswordHash string
	Email        string
}

/*
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
*/
func createChallengeDone(db *sql.DB, username string, challenge int) error {
	userID, err := queries.GetUserId(ctx, username)
	puserID := int32ToInt4(userID)

	queries.CreateChallenge(ctx, kaggaoc.CreateChallengeParams{
		Challenge: int32(challenge),
		UserID:    puserID})
	_, err = db.Exec("INSERT INTO completed (challenge, user_id) VALUES ($1, $2)", challenge, userID)
	if err != nil {
		return err
	}

	return nil
}
func getCompletedChallenges(db *sql.DB, username string) (error, []int32) {
	var completed []int32
	userID, err := queries.GetUserId(ctx, username)
	puserID := int32ToInt4(userID)

	completed, err = queries.GetChallengeCompleted(ctx, puserID)
	if err != nil {
		return err, nil
	}
	return nil, completed

}
