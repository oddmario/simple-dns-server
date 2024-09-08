# Stage 1: Build the Go application
FROM golang:1.22.6-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o simpledns .

# Stage 2: Create a minimal image with the built application
FROM gcr.io/distroless/static:latest

COPY --from=builder /app/simpledns /simpledns
ENTRYPOINT ["/simpledns"]

EXPOSE 53