# Clyde Novus

📦 [Documentation](#) | 🌿 [Node.js Client SDK (NPM)](#) | 📃 [REST API Reference](#)

---

**_CAUTION:_** Novus is not yet fully stable and is still being tested against many web vulnerabilities.

One way we're measuring progress in this regard is by going through the OWASP complete list of web vulnerabilities. You can see our progress [here](#).

---

Novus is a simple, framework-agnostic, Go-extensible auth server that has everything you need to secure your applications:

- Credentials provider _(with OAuth2 and WebAuthn arriving in the next major release, v2)_
- Two-factor authentication: TOTPs via Authenticator Apps, Email, and SMS
- Stateless session management with JSON Web Tokens (JWT).
- `HttpOnly`, `Secure`, `Samesite`, `Path=/` Cookies to maximize protection against web vulnerabilities such as Cross-Site Scripting (XSS).
- Ratelimiting and throttling utilities with an in-memory implementation (for servers) and a key-value store implementation (for serverless).
- Node.js SDK to seamlessly integrate Novus into your frontend applications with guarantees for type-safe HTTP calls.
  - Alternatively, you can also import our entire Node.js SDK into 1 file.

## Usage

### As a dedicated auth microservice

WIP.

### As a starter for monolithic backends

WIP.

### As a single-sign on (SSO) server

WIP.

## Learn more

We've written the following resources to help you integrate Novus into your own applications:

- [Novus Documentation](#)
- [Novus REST API Reference](#)

## Roadmap for Version 2.0.0

The following are the biggest features that will be added into Novus in its next major release:

- _Coming in v1.1.0_ - OAuth (Google, GitHub, Facebook) and WebAuthn (Passkeys) as alternate providers
- _Coming in v1.2.0_ - Key-value (KV) store adapter for ratelimiting and request throttling support in serverless environments
- _Coming in v1.3.0_ - Single Sign-On (SSO) mode
- _Coming in v1.4.0_ - Client SDKs for Go, Python, Java

## License

Novus is GPL-3.0 (General Public License) licensed.

At Clyde, we are committed to open-sourcing the entire distribution of our software. Please keep any forks and/or hard clones open-source.
