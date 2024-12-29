-- name: GetUserByLogin :one
SELECT * FROM users WHERE login = $1;

-- name: CreateUser :one
INSERT INTO users (login, pwd_hash, created_at) VALUES ($1, $2, $3) RETURNING *;
