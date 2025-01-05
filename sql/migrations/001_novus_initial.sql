-- +goose Up

CREATE TABLE IF NOT EXISTS clyde_user (
    id TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    recovery_code TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) STRICT;

CREATE TABLE IF NOT EXISTS clyde_email_verification_request (
    user_id TEXT NOT NULL UNIQUE PRIMARY KEY REFERENCES clyde_user(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    code TEXT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS clyde_email_update_request (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES clyde_user(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL,
    code TEXT NOT NULL
) STRICT;

CREATE INDEX IF NOT EXISTS clyde_email_update_request_user_id_index ON clyde_email_update_request(user_id);

CREATE TABLE IF NOT EXISTS clyde_password_reset_request (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES clyde_user(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    code_hash TEXT NOT NULL
) STRICT;

CREATE INDEX IF NOT EXISTS clyde_password_reset_request_user_id_index ON clyde_password_reset_request(user_id);

CREATE TABLE IF NOT EXISTS clyde_user_totp_credential (
    user_id TEXT NOT NULL PRIMARY KEY REFERENCES clyde_user(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    key BLOB NULL
) STRICT;

CREATE TABLE IF NOT EXISTS clyde_passkey_credential (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL REFERENCES clyde_user(id),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    cose_algorithm_id INTEGER NOT NULL,
    public_key BLOB NULL
) STRICT;

CREATE INDEX IF NOT EXISTS clyde_passkey_credential_user_id_index ON clyde_passkey_credential(user_id);

CREATE TABLE IF NOT EXISTS clyde_security_key (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL REFERENCES clyde_user(id),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    cose_algorithm_id INTEGER NOT NULL,
    public_key BLOB NULL
) STRICT;

CREATE INDEX IF NOT EXISTS clyde_security_key_user_id_index ON clyde_security_key(user_id);

-- +goose Down

DROP INDEX IF EXISTS clyde_security_key_user_id_index;
DROP TABLE IF EXISTS clyde_security_key;

DROP INDEX IF EXISTS clyde_passkey_credential_user_id_index;
DROP TABLE IF EXISTS clyde_passkey_credential;

DROP TABLE IF EXISTS clyde_user_totp_credential;

DROP INDEX IF EXISTS clyde_password_reset_request_user_id_index;
DROP TABLE IF EXISTS clyde_password_reset_request;

DROP INDEX IF EXISTS clyde_email_update_request_user_id_index;
DROP TABLE IF EXISTS clyde_email_update_request;

DROP TABLE IF EXISTS clyde_email_verification_request;

DROP TABLE IF EXISTS clyde_user;