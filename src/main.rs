use axum::Router;
use routes::{create_auth_routes, create_session_routes, create_user_routes};
use serde::Deserialize;
use sqlx::postgres::PgPoolOptions;

mod dtos;
mod entities;
mod handlers;
mod routes;

#[derive(Deserialize, Debug)]
struct EnvironmentVariables {
    database_url: String,
}

#[tokio::main]
async fn main() {
    // Load environment variable
    dotenvy::dotenv().ok();
    let env = envy::from_env::<EnvironmentVariables>().unwrap();

    // Create a database connection pool
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(env.database_url.as_str())
        .await
        .unwrap();

    // Auto-migrate
    if let Err(e) = sqlx::migrate!("./migrations").run(&pool).await {
        eprintln!("Migration error: {:?}", e);
        std::process::exit(1);
    }

    // Initialize API routes
    let routes = Router::new()
        .merge(create_auth_routes())
        .merge(create_user_routes())
        .merge(create_session_routes());

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();
    axum::serve(listener, Router::new().nest("/api", routes))
        .await
        .unwrap();
}
