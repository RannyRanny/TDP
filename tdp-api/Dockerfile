# Используйте официальный образ Go как базовый
FROM golang:1.17 as builder

# Установите рабочий каталог в контейнере
WORKDIR /app

# Копируйте модули зависимостей
COPY go.mod .
COPY go.sum .

# Загрузите зависимости
RUN go mod download

# Копируйте исходный код проекта
COPY . .

# Соберите исполняемый файл для нашего приложения.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task-manager ./cmd/main.go

# Начните новый этап с небольшого образа
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Скопируйте исполняемый файл из предыдущего этапа
COPY --from=builder /app/task-manager .

# Объявите порт, который будет использовать приложение
EXPOSE 8080

# Запустите исполняемый файл
CMD ["./tdp-api"]
