-- name: FindUserById :one
SELECT * FROM thorfinn_users WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM thorfinn_users WHERE email = $1;


-- name: ListUsers :many
SELECT id, email, verified, two_factor_enabled, created_at, updated_at
FROM thorfinn_users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO thorfinn_users (id, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUserVerified :one
UPDATE thorfinn_users
SET verified = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE thorfinn_users
SET password_hash = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUser :one
UPDATE thorfinn_users
SET email = $2, password_hash = $3, verified = $4, two_factor_enabled = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM thorfinn_users WHERE id = $1;
