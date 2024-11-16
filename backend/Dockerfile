# FROM golang:latest

# RUN go install github.com/air-verse/air@latest

# WORKDIR /careerscrape/backend

# COPY go.mod go.sum ./


# COPY . .

# RUN go mod download

# # RUN go build -o main
# EXPOSE 8000


# Stage 1: Build
FROM golang:latest AS builder

WORKDIR /careerscrape/backend
COPY go.mod go.sum ./
RUN go mod download
COPY .env .
RUN chmod 644 .env

COPY . .


# Build a static binary to avoid dependency issues
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /careerscrape/backend/main .
COPY --from=builder /careerscrape/backend/.env .


# Ensure the binary is executable
RUN chmod +x /root/main

EXPOSE 8000

CMD ["./main"]
