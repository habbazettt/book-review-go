FROM golang:1.24.0-alpine3.20 AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Final image
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
