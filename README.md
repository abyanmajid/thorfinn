# Thorfinn

**Thorfinn** is a simple, framework-agnostic authentication server built with the [Matcha framework](https://github.com/abyanmajid/matcha). It provides a complete credentials + email OTP authentication solution.

You can use **Thorfinn** as a standalone authentication microservice, or as a starter for your RESTful backend.

## Development

Copy the `.env.example` file to `.env`:

```
cp .env.example .env
```

Serve local postgres (port 5433) and mailhog (port 8026) containers:

```
docker-compose up -d
```

Apply migrations (NOTE: To down-migrate, run `make migrate-down`):

```
make migrate-up
```

Serve the application on port 8080:

```
make dev
```

## Production

The server requires the following environment variables:

- `ORIGIN`: The URL of the server.
- `FRONTEND_URL`: The URL of the frontend.
- `DATABASE_URL`: The database URL.
- `SMTP_HOST`: The SMTP host.
- `SMTP_PORT`: The SMTP port.
- `SMTP_USER`: The SMTP user.
- `SMTP_PASSWORD`: The SMTP password.
- `EMAIL_FROM`: The email address of the sender.
- `JWT_SECRET`: The JWT secret key.
- `ENCRYPTION_SECRET`: The encryption secret key.

It's advised to serve the production server using Docker. To build the docker image, run:

```
docker build -t thorfinn .
```

To run the docker container, run:

```
docker run -d -p 8080:8080 thorfinn
```

## License

Thorfinn is licensed under GPL 3.0