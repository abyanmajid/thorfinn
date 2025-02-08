# Thorfinn

**Thorfinn** is a simple, framework-agnostic, Go-extensible auth server that provides a complete credentials + email OTP authentication solution.

## Development

Copy the `.env.example` file to `.env`:

```
cp .env.example .env
```

Serve local postgres (port 5433) and mailhog (port 8026) containers:

```
docker-compose up -d
```

Serve the application on port 8080:

```
make dev
```

## Production

The server requires the following environment variables:

- `DEBUG`: Whether to run in debug mode.
- `ORIGIN`: The origin of the server.
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