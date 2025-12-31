-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateApiKey :one
INSERT INTO api_keys (user_id, name, key_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetApiKeysByUser :many
SELECT id, name, last_used_at, revoked_at, created_at
FROM api_keys
WHERE user_id = $1;

-- name: GetApiKeyByHash :one
SELECT * From api_keys
WHERE key_hash = $1 AND revoked_at IS NULL;

-- name: RevokeApiKey :exec
UPDATE api_keys
SET revoked_at = CURRENT_TIMESTAMP
WHERE id = $1
AND user_id = $2
AND revoked_at IS NULL;
