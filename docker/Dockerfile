FROM golang:alpine AS builder

WORKDIR /app

COPY go.* .

RUN go mod tidy

COPY . .

RUN go build -o bin/api cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/api .

CMD ["./api"]