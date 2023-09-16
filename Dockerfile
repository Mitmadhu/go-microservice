FROM golang:1.21.1 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/main

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service

# RUN go build -o user-service

###############

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main/user-service .
COPY static /app/static

EXPOSE 3000

CMD ["./user-service"]
