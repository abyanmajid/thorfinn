-- name: FindEmailUpdateRequestById :one
SELECT * FROM clyde_email_update_request
WHERE id = $1;

-- name: ListEmailUpdateRequests :many
SELECT * FROM clyde_email_update_request
ORDER BY created_at ASC;

-- name: InsertEmailUpdateRequest :exec
INSERT INTO clyde_email_update_request (id, user_id, expires_at, email, code)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateEmailUpdateRequest :exec
UPDATE clyde_email_update_request
SET updated_at = CURRENT_TIMESTAMP,
    expires_at = COALESCE($2, expires_at),
    email = COALESCE($3, email),
    code = COALESCE($4, code)
WHERE id = $1;

-- name: DeleteEmailUpdateRequest :exec
DELETE FROM clyde_email_update_request
WHERE id = $1;
