FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY video/go.mod video/go.sum ./video/
COPY database ./database

WORKDIR /app/video
RUN go mod download

COPY video/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /video .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /video /video

EXPOSE 8084
ENTRYPOINT ["/video"]