FROM golang:1.25.3-alpine3.21 AS base

FROM base AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /backend ./cmd/main.go

FROM base AS final
RUN apk add --no-cache tzdata
COPY --from=builder /backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]

