FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git build-base curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/main .

FROM alpine:3.18 AS runtime

RUN apk --no-cache add bash curl postgresql-client

WORKDIR /app

COPY --from=builder /app/build/main /app/main

COPY .env .env

EXPOSE ${PORT}

CMD ["/app/main"]
