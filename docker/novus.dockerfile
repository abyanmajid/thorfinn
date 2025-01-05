####################################################################################################
## Build Stage
####################################################################################################
FROM rust:latest AS builder

RUN rustup target add x86_64-unknown-linux-musl
RUN apt update && apt install -y musl-tools musl-dev libssl-dev pkg-config
RUN update-ca-certificates

ENV USER=clyde
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /clyde-novus

COPY . .

# Set OPENSSL_DIR for static linking (this ensures the correct OpenSSL is found)
ENV OPENSSL_DIR=/usr/include/openssl

RUN cargo build --release --target x86_64-unknown-linux-musl

####################################################################################################
## Runtime Stage
####################################################################################################
FROM alpine:latest

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /clyde-novus

COPY --from=builder /clyde-novus/target/x86_64-unknown-linux-musl/release/clyde-novus ./

USER clyde:clyde

CMD ["/clyde-novus/clyde-novus"]
