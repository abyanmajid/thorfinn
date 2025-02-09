-- name: CreateBlacklistedToken :one
INSERT INTO thorfinn_blacklisted_tokens (id, token) VALUES ($1, $2) RETURNING *;

-- name: GetBlacklistedToken :one
SELECT * FROM thorfinn_blacklisted_tokens WHERE token = $1;
