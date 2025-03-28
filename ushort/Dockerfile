FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ushort .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/ushort /ushort

COPY .env .

RUN chmod +x /ushort

CMD ["/ushort"]

EXPOSE 8080