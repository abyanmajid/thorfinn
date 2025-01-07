### Step 1: Build stage
FROM golang:1.22.4-alpine as builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o clyde-orion ./cmd

### 
## Step 2: Runtime stage
FROM alpine:3.20

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/clyde-orion /

# Set the entry point for the container
ENTRYPOINT ["/clyde-orion"]