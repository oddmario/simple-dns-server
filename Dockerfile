# Step 1: Builder stage
FROM golang:1.22.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code and build the application
COPY . .
RUN go build -o simpledns -ldflags="-w -s" -trimpath -v .

# Step 2: Runner stage
FROM gcr.io/distroless/static:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/simpledns /simpledns

# Set the entrypoint to run the application
ENTRYPOINT ["/simpledns"]

# Expose port 53 for Docker Compose usage
EXPOSE 53