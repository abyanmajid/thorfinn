-- name: FindPasskeyCredentialById :one
SELECT * FROM clyde_passkey_credential
WHERE id = $1;

-- name: ListPasskeyCredentials :many
SELECT * FROM clyde_passkey_credential
ORDER BY created_at ASC;

-- name: InsertPasskeyCredential :exec
INSERT INTO clyde_passkey_credential (id, user_id, name, cose_algorithm_id, public_key)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdatePasskeyCredential :exec
UPDATE clyde_passkey_credential
SET updated_at = CURRENT_TIMESTAMP,
    name = COALESCE($2, name),
    cose_algorithm_id = COALESCE($3, cose_algorithm_id),
    public_key = COALESCE($4, public_key)
WHERE id = $1;

-- name: DeletePasskeyCredential :exec
DELETE FROM clyde_passkey_credential
WHERE id = $1;