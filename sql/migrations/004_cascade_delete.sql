-- +goose Up

ALTER TABLE clyde_email_verification_request
    DROP CONSTRAINT clyde_email_verification_request_user_id_fkey,
    ADD CONSTRAINT clyde_email_verification_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

ALTER TABLE clyde_email_update_request
    DROP CONSTRAINT clyde_email_update_request_user_id_fkey,
    ADD CONSTRAINT clyde_email_update_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

ALTER TABLE clyde_password_reset_request
    DROP CONSTRAINT clyde_password_reset_request_user_id_fkey,
    ADD CONSTRAINT clyde_password_reset_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

ALTER TABLE clyde_user_totp_credential
    DROP CONSTRAINT clyde_user_totp_credential_user_id_fkey,
    ADD CONSTRAINT clyde_user_totp_credential_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

ALTER TABLE clyde_passkey_credential
    DROP CONSTRAINT clyde_passkey_credential_user_id_fkey,
    ADD CONSTRAINT clyde_passkey_credential_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

ALTER TABLE clyde_security_key
    DROP CONSTRAINT clyde_security_key_user_id_fkey,
    ADD CONSTRAINT clyde_security_key_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id) ON DELETE CASCADE;

-- +goose Down

ALTER TABLE clyde_email_verification_request
    DROP CONSTRAINT clyde_email_verification_request_user_id_fkey,
    ADD CONSTRAINT clyde_email_verification_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

ALTER TABLE clyde_email_update_request
    DROP CONSTRAINT clyde_email_update_request_user_id_fkey,
    ADD CONSTRAINT clyde_email_update_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

ALTER TABLE clyde_password_reset_request
    DROP CONSTRAINT clyde_password_reset_request_user_id_fkey,
    ADD CONSTRAINT clyde_password_reset_request_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

ALTER TABLE clyde_user_totp_credential
    DROP CONSTRAINT clyde_user_totp_credential_user_id_fkey,
    ADD CONSTRAINT clyde_user_totp_credential_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

ALTER TABLE clyde_passkey_credential
    DROP CONSTRAINT clyde_passkey_credential_user_id_fkey,
    ADD CONSTRAINT clyde_passkey_credential_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

ALTER TABLE clyde_security_key
    DROP CONSTRAINT clyde_security_key_user_id_fkey,
    ADD CONSTRAINT clyde_security_key_user_id_fkey FOREIGN KEY (user_id) REFERENCES clyde_user(id);

