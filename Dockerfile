FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./backend ./backend

RUN go build -o ./backend/chess-backend ./backend 

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/backend/chess-backend .
COPY --from=builder /app/backend/.env .

EXPOSE 8080

CMD ["./chess-backend"]
