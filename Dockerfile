#build
FROM golang:1.13-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o single-srv

# run
FROM alpine:3.11

COPY --from=builder /app/single-srv /

EXPOSE 8000

ENTRYPOINT ["/single-srv"]
