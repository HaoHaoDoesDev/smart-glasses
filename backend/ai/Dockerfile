FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY ai/go.mod ai/go.sum ./ai/
COPY database ./database

WORKDIR /app/ai
RUN go mod download

COPY ai/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /ai .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /ai /ai

EXPOSE 8086
ENTRYPOINT ["/ai"]