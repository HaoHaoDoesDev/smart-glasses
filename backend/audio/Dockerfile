FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY audio/go.mod audio/go.sum ./audio/
COPY database ./database

WORKDIR /app/audio
RUN go mod download

COPY audio/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /audio .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /audio /audio

EXPOSE 8085
ENTRYPOINT ["/audio"]