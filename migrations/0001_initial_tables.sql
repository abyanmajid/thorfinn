CREATE TABLE IF NOT EXISTS nova_user (
    id TEXT NOT NULL PRIMARY KEY,
    created_at INTEGER NOT NULL,
    password_hash TEXT NOT NULL,
    recovery_code TEXT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS nova_user_email_verification_request (
    user_id TEXT NOT NULL UNIQUE PRIMARY KEY REFERENCES nova_user(id),
    created_at INTEGER NOT NULL,
    expires_at INTEGER NOT NULL,
    code TEXT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS nova_email_update_request (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES nova_user(id),
    created_at INTEGER NOT NULL,
    expires_at INTEGER NOT NULL,
    email TEXT NOT NULL,
    code TEXT NOT NULL
) STRICT;

CREATE INDEX IF NOT EXISTS nova_email_update_request_user_id_index ON nova_email_update_request(user_id);

CREATE TABLE IF NOT EXISTS nova_password_reset_request (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES nova_user(id),
    created_at INTEGER NOT NULL,
    expires_at INTEGER NOT NULL,
    code_hash TEXT NOT NULL
) STRICT;

CREATE INDEX IF NOT EXISTS nova_password_reset_request_user_id_index ON nova_password_reset_request(user_id);

CREATE TABLE IF NOT EXISTS nova_user_totp_credential (
    user_id TEXT NOT NULL PRIMARY KEY REFERENCES nova_user(id),
    created_at INTEGER NOT NULL,
    key BLOB NULL
) STRICT;

CREATE TABLE IF NOT EXISTS nova_passkey_credential (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL REFERENCES nova_user(id),
    name TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    cose_algorithm_id INTEGER NOT NULL,
    public_key BLOB NULL
) STRICT;

CREATE INDEX IF NOT EXISTS nova_passkey_credential_user_id_index ON nova_passkey_credential(user_id);

CREATE TABLE IF NOT EXISTS nova_security_key (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL REFERENCES nova_user(id),
    name TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    cose_algorithm_id INTEGER NOT NULL,
    public_key BLOB NULL
) STRICT;

CREATE INDEX IF NOT EXISTS nova_security_key_user_id_index ON nova_security_key(user_id);
