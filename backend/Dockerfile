FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./chess-backend ./

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/chess-backend ./
COPY --from=builder /app/.env ./

EXPOSE 8080

CMD ["./chess-backend"]
