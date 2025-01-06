-- +goose Up

ALTER TABLE clyde_user
ADD COLUMN is_email_verified BOOLEAN DEFAULT FALSE;

ALTER TABLE clyde_user
ADD COLUMN is_2fa_enabled BOOLEAN DEFAULT FALSE;

-- +goose Down

ALTER TABLE clyde_user
DROP COLUMN IF EXISTS is_email_verified;

ALTER TABLE clyde_user
DROP COLUMN IF EXISTS is_2fa_enabled;
