FROM golang:1.23.4 AS builder

WORKDIR /service3/

COPY . .

RUN CGO_ENABLED=0 go build -o microservice /service3/main.go

FROM alpine:latest

WORKDIR /service3

COPY --from=builder /service3/microservice /service3/microservice

EXPOSE 3003

CMD ["./microservice"]
