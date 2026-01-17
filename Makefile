# Makefile for Web Wake-on-LAN

# Get version from git
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Linker flags
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

# Binary name
BINARY_NAME := web-wol

.PHONY: all build build-frontend build-backend clean dev

all: build

# Build everything
build: build-frontend build-backend

# Build frontend assets
build-frontend:
	cd frontend && pnpm install && pnpm build

# Build Go binary with version injection
build-backend:
	go build $(LDFLAGS) -o $(BINARY_NAME) main.go

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -rf frontend/dist

# Run development server (backend)
dev:
	go run main.go
