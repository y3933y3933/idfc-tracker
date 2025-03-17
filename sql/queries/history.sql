-- name: InsertHistory :exec
INSERT INTO history (user_id, point, reason)
VALUES (?, ?, ?);

-- name: ClearUserHistory :exec
DELETE FROM history WHERE user_id = ?;