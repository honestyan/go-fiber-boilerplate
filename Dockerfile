FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git build-base curl

RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b /usr/local/bin

ENV PATH="/usr/local/bin:$PATH"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.18 AS runtime

RUN apk --no-cache add bash curl postgresql-client

WORKDIR /app

COPY --from=builder /app/main /app

EXPOSE ${PORT}

CMD ["/app/main"]
