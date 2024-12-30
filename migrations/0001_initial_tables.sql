CREATE TYPE user_role AS ENUM ('user', 'admin');
CREATE TYPE auth_method AS ENUM ('password', 'oauth', 'webauthn');

-- Users table for storing primary user information
CREATE TABLE openyan_users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(256), -- Nullable for OAuth users
    role user_role NOT NULL DEFAULT 'user',
    is_banned BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT email_or_password CHECK (
        (email IS NOT NULL AND password IS NOT NULL) OR 
        (email IS NOT NULL AND password IS NULL) -- Allow OAuth with email only
    )
);

CREATE UNIQUE INDEX openyan_users_email_idx ON openyan_users (email);

-- Authentication methods, e.g., for OAuth, 2FA, or WebAuthn
CREATE TABLE openyan_auth_methods (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id UUID REFERENCES openyan_users(id) ON DELETE CASCADE,
    method auth_method NOT NULL,
    provider VARCHAR(128), -- For OAuth: 'google', 'github', etc.
    provider_id VARCHAR(256), -- External provider user ID
    secret VARCHAR(512), -- For TOTP secrets, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE(user_id, method, provider) -- Prevent duplicate methods per user
);

-- Two-factor authentication tokens (OTP for email, SMS, etc.)
CREATE TABLE openyan_2fa_tokens (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id UUID REFERENCES openyan_users(id) ON DELETE CASCADE,
    token VARCHAR(6) NOT NULL, -- OTP code
    method auth_method NOT NULL, -- 'authenticator', 'email', 'sms'
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Sessions table for managing JWT or database sessions
CREATE TABLE openyan_sessions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id UUID REFERENCES openyan_users(id) ON DELETE CASCADE,
    user_agent TEXT,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
);