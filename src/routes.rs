use axum::{
    routing::{delete, get, post, put},
    Router,
};

use crate::handlers::register_handler;

pub fn create_auth_routes() -> Router {
    Router::new()
        .route("/auth/credentials/register", post(register_handler))
        .route("/auth/credentials/login", post(register_handler))
        .route("/auth/oauth/:provider/login", post(register_handler))
        .route("/auth/oauth/:provider/callback", post(register_handler))
        .route("/auth/webauthn/initiate", post(register_handler))
        .route("/auth/webauthn/verify", post(register_handler))
        .route("/auth/logout", post(register_handler))
        .route("/auth/2fa/authenticator/initiate", post(register_handler))
        .route("/auth/2fa/authenticator/verify", post(register_handler))
        .route("/auth/2fa/email/initiate", post(register_handler))
        .route("/auth/2fa/email/verify", post(register_handler))
        .route("/auth/2fa/sms/initiate", post(register_handler))
        .route("/auth/2fa/sms/verify", post(register_handler))
}

pub fn create_user_routes() -> Router {
    Router::new()
        .route("/user/me", get(register_handler))
        .route("/user/:id", get(register_handler))
        .route("/user/:id", put(register_handler))
        .route("/user/:id", delete(register_handler))
}

pub fn create_session_routes() -> Router {
    Router::new()
        .route("/sessions", get(register_handler))
        .route("/sessions", post(register_handler))
        .route("/sessions/refresh", post(register_handler))
        .route("/sessions/revoke", post(register_handler))
}
