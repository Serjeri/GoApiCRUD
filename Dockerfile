# Этап сборки (используем актуальную версию Go)
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем только файлы зависимостей для кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./main.go

# Этап запуска
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env
# Устанавливаем tzdata для работы с временными зонами
RUN apk add --no-cache tzdata


# Команда запуска
CMD ["/app/main"]