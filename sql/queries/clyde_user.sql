-- name: FindUserById :one
SELECT * FROM clyde_user
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM clyde_user
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM clyde_user
ORDER BY created_at ASC;

-- name: InsertUser :exec
INSERT INTO clyde_user (id, email, password_hash, recovery_code, name, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, COALESCE($6, 'user'), COALESCE($7, CURRENT_TIMESTAMP), COALESCE($8, CURRENT_TIMESTAMP));

-- name: UpdateUser :exec
UPDATE clyde_user
SET email = COALESCE($2, email),
    password_hash = COALESCE($3, password_hash),
    recovery_code = COALESCE($4, recovery_code),
    name = COALESCE($5, name),
    role = COALESCE($6, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM clyde_user
WHERE id = $1;
