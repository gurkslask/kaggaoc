// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package kaggaoc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Completed struct {
	CompleteID int32
	UserID     pgtype.Int4
	Challenge  int32
}

type User struct {
	UserID       int32
	Username     string
	PasswordHash string
	Email        string
	Seed         string
}
