# Clyde Novus

📦 [Documentation](#) | 🌿 [Node.js Client SDK (NPM)](#) | 📃 [REST API Reference](#)

---

**_CAUTION:_** Novus is not yet fully stable and is still being tested against many web vulnerabilities.

One way we're measuring progress in this regard is by going through the OWASP complete list of web vulnerabilities. You can see our progress [here](#).

---

Novus is a simple, framework-agnostic, Go-extensible auth server that has everything you need to secure your applications:

- Multiple providers: Credentials, OAuth2, WebAuthn (Passkeys).
- Two-factor authentication: TOTPs via Authenticator Apps, Email, and SMS
- Stateless session management with JSON Web Tokens (JWT).
- `HttpOnly`, `Secure`, `Samesite`, `Path=/` Cookies to maximize protection against web vulnerabilities such as Cross-Site Scripting (XSS).
- Ratelimiting and throttling utilities with an in-memory implementation (for servers) and a key-value store implementation (for serverless).
- Node.js SDK to seamlessly integrate Novus into your frontend applications with guarantees for type-safe HTTP calls.
  - Alternatively, you can also import our entire Node.js SDK into 1 file.

## Ways to Use Novus

### As an Auth Microservice

WIP.

### As a Single Sign-On (SSO)

WIP.

### As a Backend Starter for Monoliths

WIP.

## Learn More

WIP.

## License

Novus is GPL-3.0 (General Public License) licensed.
