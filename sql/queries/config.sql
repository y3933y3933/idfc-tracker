-- name: GetActiveUserID :one
SELECT value 
FROM config 
WHERE key = 'active_user_id';

-- name: SetActiveUserID :exec
INSERT INTO config (key, value) 
VALUES ('active_user_id', ?) 
ON CONFLICT(key) DO UPDATE SET value = excluded.value;