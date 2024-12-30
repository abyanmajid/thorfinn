use axum::{http::StatusCode, Json};

use crate::dtos::RegisterResponsePayload;

pub async fn register_handler() -> (StatusCode, Json<RegisterResponsePayload>) {
    (
        StatusCode::CREATED,
        Json(RegisterResponsePayload {
            user_id: String::from(""),
        }),
    )
}
