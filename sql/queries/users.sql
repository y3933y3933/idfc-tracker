-- name: GetUserByName :one
SELECT id, name 
FROM users 
WHERE name = ?;

-- name: CreateUser :exec
INSERT INTO users (name) VALUES (?);

