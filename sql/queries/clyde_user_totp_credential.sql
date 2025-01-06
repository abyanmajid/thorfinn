-- name: FindTotpCredentialByUserId :one
SELECT * FROM clyde_user_totp_credential
WHERE user_id = $1;

-- name: ListTotpCredentials :many
SELECT * FROM clyde_user_totp_credential
ORDER BY created_at ASC;

-- name: InsertTotpCredential :exec
INSERT INTO clyde_user_totp_credential (user_id, key)
VALUES ($1, $2);

-- name: UpdateTotpCredential :exec
UPDATE clyde_user_totp_credential
SET updated_at = CURRENT_TIMESTAMP,
    key = COALESCE($2, key)
WHERE user_id = $1;

-- name: DeleteTotpCredential :exec
DELETE FROM clyde_user_totp_credential
WHERE user_id = $1;
