-- name: CreatePoint :exec
INSERT INTO points (user_id, goal, created_at)
VALUES (?, ?, CURRENT_TIMESTAMP);