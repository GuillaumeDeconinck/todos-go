FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o api ./cmd/api.go

# Create the minimal runtime image
FROM alpine AS prod

WORKDIR /app

COPY --from=builder /app/api ./api

CMD ["/app/api"]
