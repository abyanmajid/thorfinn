-- name: FindUserById :one
SELECT id, name, email, role, 
       is_email_verified, is_2fa_enabled, 
       created_at, updated_at
FROM clyde_user
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT id, name, email, role, 
       is_email_verified, is_2fa_enabled, 
       created_at, updated_at
FROM clyde_user
WHERE email = $1;

-- name: ListUsers :many
SELECT id, name, email, role, 
       is_email_verified, is_2fa_enabled, 
       created_at, updated_at
FROM clyde_user
ORDER BY created_at ASC;

-- name: InsertUser :one
INSERT INTO clyde_user (id, email, password_hash, recovery_code, name, role)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateUser :one
UPDATE clyde_user
SET email = COALESCE($2, email),
    password_hash = COALESCE($3, password_hash),
    recovery_code = COALESCE($4, recovery_code),
    name = COALESCE($5, name),
    role = COALESCE($6, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM clyde_user
WHERE id = $1
RETURNING *;
