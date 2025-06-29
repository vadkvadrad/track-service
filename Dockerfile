FROM golang:1.23.10 AS builder

WORKDIR /app
COPY go.mod .
COPY .env /app/.env
COPY . .
RUN CGO_ENABLED=0 go build -o /track-service cmd/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/go.mod /app/go.mod
COPY --from=builder /track-service /track-service
CMD ["/track-service"]