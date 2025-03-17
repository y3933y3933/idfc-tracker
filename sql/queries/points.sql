-- name: CreatePoint :exec
INSERT INTO points (user_id, goal, created_at)
VALUES (?, ?, CURRENT_TIMESTAMP);

-- name: GetPointByUserID :one
SELECT id, user_id, total, goal 
FROM points
WHERE user_id = ?;

-- name: UpdateTotalByUserID :exec
UPDATE points
SET total = ?
WHERE user_id = ?;

-- name: UpdateGoalByUserID :exec
UPDATE points
SET goal = ?
WHERE user_id = ?;