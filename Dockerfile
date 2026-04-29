# Stage 1: Build Frontend
FROM node:22-alpine AS frontend-builder
WORKDIR /app/frontend

# Install pnpm
RUN npm install -g pnpm

# Copy package files and install dependencies
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# Copy frontend source and build
COPY frontend/ ./
RUN pnpm build

# Stage 2: Build Backend
FROM golang:1.24-alpine AS backend-builder
ARG VERSION=dev
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy all source code
COPY . .

# Copy built frontend assets from the previous stage
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist

# Build the Go binary with version injection
RUN go build -ldflags "-X main.Version=${VERSION}" -o web-wol main.go

# Stage 3: Runtime
FROM alpine:3.21

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the binary from the backend-builder stage
COPY --from=backend-builder /app/web-wol .

# Expose the application port
EXPOSE 8080

# Start the application
ENTRYPOINT ["/app/web-wol"]
