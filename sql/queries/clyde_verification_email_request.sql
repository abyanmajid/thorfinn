-- name: FindEmailVerificationRequestByUserId :one
SELECT * FROM clyde_email_verification_request
WHERE user_id = $1;

-- name: ListEmailVerificationRequests :many
SELECT * FROM clyde_email_verification_request
ORDER BY created_at ASC;

-- name: InsertEmailVerificationRequest :exec
INSERT INTO clyde_email_verification_request (user_id, expires_at, code)
VALUES ($1, $2, $3);

-- name: UpdateEmailVerificationRequest :exec
UPDATE clyde_email_verification_request
SET updated_at = CURRENT_TIMESTAMP,
    expires_at = COALESCE($2, expires_at),
    code = COALESCE($3, code)
WHERE user_id = $1;

-- name: DeleteEmailVerificationRequest :exec
DELETE FROM clyde_email_verification_request
WHERE user_id = $1;

-- name: DeleteExpiredEmailVerificationRequests :exec
DELETE FROM clyde_email_verification_request
WHERE expires_at <= $1;