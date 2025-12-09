# Этап сборки (builder)
FROM golang:1.25 AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o pz10-auth ./cmd/server

# Этап рантайма
FROM alpine:3.20

WORKDIR /app

# Обновляем сертификаты и часовые пояса (по необходимости)
RUN apk add --no-cache ca-certificates tzdata

# Копируем собранный бинарник из builder-образа
COPY --from=builder /app/pz10-auth ./pz10-auth

# Порт приложения внутри контейнера
EXPOSE 8080

# Команда запуска
CMD ["./pz10-auth"]
