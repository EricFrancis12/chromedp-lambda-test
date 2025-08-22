FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./
RUN go mod download
RUN go build -o main

FROM chromedp/headless-shell:113.0.5672.93

WORKDIR /app

# Install CA certificates to allow the container to verify SSL/TLS connections to external services.
# All http requests will fail if not done.
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]
