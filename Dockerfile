FROM golang:1.21-alpine AS builder

ARG PORT=8080
ARG APP_ENV=development

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o payment-gw

FROM alpine:3.18
ARG PORT
ARG APP_ENV
ENV PORT=$PORT
ENV APP_ENV=$APP_ENV

WORKDIR /app
COPY --from=builder /app/payment-gw .
COPY --from=builder /app/config/config.${APP_ENV}.ini ./config/config.${APP_ENV}.ini
COPY --from=builder /app/migrations ./migrations

EXPOSE $PORT
CMD ["./payment-gw", "serve"]