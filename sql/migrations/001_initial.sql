-- +goose Up

CREATE TABLE IF NOT EXISTS thorfinn_users (
    id TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT password_min_length CHECK (length(password_hash) >= 8)
);

CREATE INDEX IF NOT EXISTS idx_thorfinn_users_email ON thorfinn_users(email);

CREATE TABLE IF NOT EXISTS thorfinn_otp_codes (
    id TEXT NOT NULL PRIMARY KEY,
    code TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT code_min_length CHECK (length(code) = 6),
    CONSTRAINT code_is_alphanumeric CHECK (code ~ '^[A-Z0-9]{6}$')
);

CREATE TABLE IF NOT EXISTS thorfinn_blacklisted_tokens (
    id TEXT NOT NULL PRIMARY KEY,
    token TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE IF EXISTS thorfinn_users;
DROP TABLE IF EXISTS thorfinn_otp_codes;
DROP TABLE IF EXISTS thorfinn_blacklisted_tokens;

DROP INDEX IF EXISTS idx_thorfinn_users_email;
