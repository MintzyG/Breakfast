FROM golang:1.26.3-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /pancake ./cmd/main.go

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /pancake .

EXPOSE 8080
CMD ["./pancake"]