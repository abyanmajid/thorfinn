use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct RegisterResponsePayload {
    pub user_id: String,
}
