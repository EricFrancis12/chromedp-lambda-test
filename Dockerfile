FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

FROM chromedp/headless-shell:113.0.5672.93

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]
