# Stage 1: Build the frontend
FROM node:24-bullseye AS frontend-builder
WORKDIR /app/client
COPY client/package*.json ./
COPY client/.twosky.json ./
RUN npm ci --quiet --no-progress
COPY client/ ./
RUN npm run build-prod

# Stage 2: Build the backend
FROM golang:1.24-bullseye AS backend-builder
ENV GOTOOLCHAIN=go1.26.4
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Fix go.mod after manual edits
RUN go mod tidy
# Copy prebuilt frontend from previous stage
COPY --from=frontend-builder /app/build ./build
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ADBlock-PD main.go

# Stage 3: Final image
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    bind9-host \
    && rm -rf /var/lib/apt/lists/*

# Set timezone
ENV TZ=Europe/Kyiv

# Create a non-privileged user
RUN useradd -u 10001 dnsuser
WORKDIR /opt/adblock-pd
RUN mkdir -p data conf && chown -R dnsuser:dnsuser /opt/adblock-pd
USER dnsuser

COPY --from=backend-builder /app/ADBlock-PD ./ADBlock-PD
COPY ADBlock-PD.yaml ./conf/ADBlock-PD.yaml

# DNS ports
EXPOSE 53/udp 53/tcp
# Admin UI ports
EXPOSE 80/tcp 3000/tcp
# DoT/DoQ port
EXPOSE 853/tcp 853/udp
# DoH port
EXPOSE 443/tcp 443/udp

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD host -W 2 google.com 127.0.0.1 || exit 1

ENTRYPOINT ["./ADBlock-PD", "--work-dir", "/opt/adblock-pd/data", "--config", "/opt/adblock-pd/conf/ADBlock-PD.yaml", "--no-permcheck"]


