FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main

FROM alpine:3 AS runner

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]
