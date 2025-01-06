-- +goose Up

ALTER TABLE clyde_user
ADD COLUMN name TEXT NOT NULL DEFAULT '',
ADD COLUMN role TEXT NOT NULL DEFAULT 'user';

CREATE INDEX IF NOT EXISTS clyde_user_email_index ON clyde_user(email);

-- +goose Down

DROP INDEX IF EXISTS clyde_user_email_index;

ALTER TABLE clyde_user
DROP COLUMN IF EXISTS name,
DROP COLUMN IF EXISTS role;
