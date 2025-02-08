### Step 1: Build stage
FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o thorfinn ./cmd

### 
## Step 2: Runtime stage
FROM alpine:3.20

COPY --from=builder /app/thorfinn /

ENTRYPOINT ["/thorfinn"]