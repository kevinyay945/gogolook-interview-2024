# Builder
FROM golang:1.22-alpine AS builder

WORKDIR /go/build

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

# Runner
FROM alpine:3.12

WORKDIR /app

# timezone
ENV TZ="Asia/Taipei"
RUN apk add --no-cache tzdata

COPY --from=builder /go/build/app ./
COPY migrations ./migrations

CMD ["./app"]
