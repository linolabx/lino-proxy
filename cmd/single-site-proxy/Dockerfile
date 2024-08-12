FROM golang:1.21.8-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o single-site-proxy ./cmd/single-site-proxy

FROM alpine:3

COPY --from=builder /app/single-site-proxy /usr/local/bin/single-site-proxy

ENTRYPOINT ["/usr/local/bin/single-site-proxy"]
