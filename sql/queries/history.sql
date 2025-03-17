-- name: InsertHistory :exec
INSERT INTO history (user_id, point, reason)
VALUES (?, ?, ?);

-- name: ClearUserHistory :exec
DELETE FROM history WHERE user_id = ?;

-- name: GetHistoryByUserID :many
SELECT id, user_id, point, reason, created_at
FROM history
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: GetHistoryByUserIDAndDateRange :many
SELECT id, user_id, point, reason, created_at
FROM history
WHERE user_id = ?
AND (created_at BETWEEN ? AND ?)
ORDER BY created_at DESC;