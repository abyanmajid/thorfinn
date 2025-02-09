<div align="center">
  <img src="https://github.com/user-attachments/assets/b0277571-4ac5-4130-8cc5-4538b897c757" width="50%" alt="matcha">
</div>

<div align="center">
  <h1>Thorfinn: Simple, Extensible Auth</h1>
</div>

<div align="center">
   A complete credentials + email OTP authentication solution built in Go with JWT and Cookies.
</div>

<br>

<div align="center">
  <a href="https://thorfinn.abydyl.net" target="_blank" rel="noopener noreferrer">
  <img
    src="https://img.shields.io/badge/DOCS-thorfinn.abydyl.net-555555?style=for-the-badge&logo=read-the-docs&logoColor=white&labelColor=333333&color=777777"
    alt="Docs"
  />
</a>
</div>

<br>

# Overview

**Thorfinn** is a simple, framework-agnostic authentication server built with the [Matcha framework](https://github.com/abyanmajid/matcha). It provides a complete credentials + email OTP authentication solution.

You can use **Thorfinn** as a standalone authentication microservice, or as a starter for your RESTful backend.

## Features

**Thorfinn** was designed to support:

- Credentials (email + password), including registration, login, logout, email verification, and password reset
- Email OTP
- Postgres User Data
- Secure, HTTP-only cookies
- JSON Web Tokens
- Token blacklisting

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

You should be able to see the server's OpenAPI specification at `/docs`, and the Scalar API Reference Client at `/reference`

## Production

The server requires the following environment variables:

- `ORIGIN`: The URL of the server.
- `ROOT_DOMAIN`: The root domain. Ideally, your frontend and backend should share the same domain (e.g., `app.com`, and `api.app.com`)
- `FRONTEND_URL`: The URL of the frontend.
- `DATABASE_URL`: The database URL.
- `SMTP_HOST`: The SMTP host.
- `SMTP_PORT`: The SMTP port.
- `SMTP_USER`: The SMTP user.
- `SMTP_PASSWORD`: The SMTP password.
- `EMAIL_FROM`: The email address of the sender.
- `JWT_SECRET`: The JWT secret key.
- `ENCRYPTION_SECRET`: The encryption secret key. This should be 32 characters long.
- `ENCRYPTION_IV`: The encryption initialization vector. This should be 6 characters long.

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
