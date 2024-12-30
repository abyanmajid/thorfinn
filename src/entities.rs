use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Serialize, Deserialize)]
pub enum UserRole {
    User,
    Admin,
}

#[derive(Debug, Serialize, Deserialize)]
pub enum AuthMethod {
    Password,
    Oauth,
    Webauthn,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct User {
    pub id: Uuid,
    pub email: String,
    pub password: Option<String>, // Nullable for OAuth users
    pub role: UserRole,
    pub is_banned: bool,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct AuthMethodRecord {
    pub id: Uuid,
    pub user_id: Uuid,
    pub method: AuthMethod,
    pub provider: Option<String>, // Optional, for OAuth: 'google', 'github', etc.
    pub provider_id: Option<String>, // External provider user ID
    pub secret: Option<String>,   // For TOTP secrets, etc.
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct TwoFactorToken {
    pub id: Uuid,
    pub user_id: Uuid,
    pub token: String,      // OTP code
    pub method: AuthMethod, // 'authenticator', 'email', 'sms'
    pub expires_at: NaiveDateTime,
    pub used: bool,
    pub created_at: NaiveDateTime,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Session {
    pub id: Uuid,
    pub user_id: Uuid,
    pub user_agent: Option<String>, // Optional user agent string
    pub expires_at: NaiveDateTime,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}
