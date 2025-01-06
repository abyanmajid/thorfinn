-- name: FindPasswordResetRequestById :one
SELECT * FROM clyde_password_reset_request
WHERE id = $1;

-- name: ListPasswordResetRequests :many
SELECT * FROM clyde_password_reset_request
ORDER BY created_at ASC;

-- name: InsertPasswordResetRequest :exec
INSERT INTO clyde_password_reset_request (id, user_id, expires_at, code_hash)
VALUES ($1, $2, $3, $4);

-- name: UpdatePasswordResetRequest :exec
UPDATE clyde_password_reset_request
SET updated_at = CURRENT_TIMESTAMP,
    expires_at = COALESCE($2, expires_at),
    code_hash = COALESCE($3, code_hash)
WHERE id = $1;

-- name: DeletePasswordResetRequest :exec
DELETE FROM clyde_password_reset_request
WHERE id = $1;

-- name: DeleteExpiredPasswordResetRequests :exec
DELETE FROM clyde_password_reset_request
WHERE expires_at <= $1;
