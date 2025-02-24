-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;
-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;
-- name: GetUserId :one
SELECT user_id FROM users
WHERE username = $1 LIMIT 1;
-- name: GetChallengeCompleted :many
SELECT challenge FROM completed
WHERE user_id = $1 ;
-- name: CreateUser :one
INSERT INTO users (
  username, password_hash, email, seed
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
-- name: CreateChallenge :one
INSERT INTO completed (
  challenge, user_id
) VALUES (
  $1, $2
)
RETURNING *;
-- name: UpdateUser :exec
UPDATE users
  set username = $2,
  email = $3
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
