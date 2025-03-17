-- name: InsertHistory :exec
INSERT INTO history (user_id, point, reason)
VALUES (?, ?, ?);