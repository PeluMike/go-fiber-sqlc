-- name: GetUsers :many
SELECT id, first_name, last_name, email, created_at::date FROM users;


-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, created_at::date FROM users
WHERE email = $1;


-- name: GetUserWithPass :one
SELECT id, first_name, last_name, email, password, created_at::date FROM users
WHERE email = $1;


-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3,$4)
RETURNING id::text, first_name, last_name, email, created_at::date; 