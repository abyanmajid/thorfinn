-- name: FindOtpCodeById :one
SELECT * FROM thorfinn_otp_codes WHERE id = $1;

-- name: CreateOtpCode :one
INSERT INTO thorfinn_otp_codes (id, code, expires_at) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteOtpCode :exec
DELETE FROM thorfinn_otp_codes WHERE id = $1;
