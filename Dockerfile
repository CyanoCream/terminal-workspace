# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go work sync && \
    go build -o bin/api ./apps/api/cmd

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/api .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/scripts/wait-for-postgres.sh ./scripts/

RUN chmod +x ./scripts/wait-for-postgres.sh && \
    apk add --no-cache postgresql-client

EXPOSE 8080

CMD ["sh", "-c", "./scripts/wait-for-postgres.sh $POSTGRES_HOST $POSTGRES_PORT && ./api"]