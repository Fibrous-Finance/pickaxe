# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o pickaxe cmd/pickaxe/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/pickaxe .
COPY app.env .
COPY app_test.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/init/states ./db/init/states
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/pickaxe" ]
ENTRYPOINT [ "/app/start.sh" ]