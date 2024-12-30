use axum::{http::StatusCode, Json};

use crate::dtos::CreateUserResponse;

pub async fn create_user() -> (StatusCode, Json<CreateUserResponse>) {
    (
        StatusCode::CREATED,
        Json(CreateUserResponse {
            user_id: String::from(""),
        }),
    )
}
