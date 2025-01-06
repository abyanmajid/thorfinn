-- name: FindSecurityKeyById :one
SELECT * FROM clyde_security_key
WHERE id = $1;

-- name: ListSecurityKeys :many
SELECT * FROM clyde_security_key
ORDER BY created_at ASC;

-- name: InsertSecurityKey :exec
INSERT INTO clyde_security_key (id, user_id, name, created_at, updated_at, cose_algorithm_id, public_key)
VALUES ($1, $2, $3, COALESCE($4, CURRENT_TIMESTAMP), COALESCE($5, CURRENT_TIMESTAMP), $6, $7);

-- name: UpdateSecurityKey :exec
UPDATE clyde_security_key
SET updated_at = CURRENT_TIMESTAMP,
    name = COALESCE($2, name),
    cose_algorithm_id = COALESCE($3, cose_algorithm_id),
    public_key = COALESCE($4, public_key)
WHERE id = $1;

-- name: DeleteSecurityKey :exec
DELETE FROM clyde_security_key
WHERE id = $1;
