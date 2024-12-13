FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app

RUN go build -o /app/airbnb

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/airbnb .

EXPOSE 8080

CMD ["./airbnb"]
