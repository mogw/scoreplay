# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /scoreplay ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

# Install bash
RUN apk add --no-cache bash

COPY --from=builder /scoreplay .

COPY wait-for-it.sh ./app/wait-for-it.sh
RUN chmod +x ./app/wait-for-it.sh


EXPOSE 8080

CMD ["./scoreplay"]
