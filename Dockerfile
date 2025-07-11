# syntax=docker/dockerfile:1.4

FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

EXPOSE 3333

CMD ["./main"]
