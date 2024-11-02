FROM golang:latest1

# Установка рабочей директории
WORKDIR /app

# Копирование всего проекта
COPY . .

# Установка зависимостей
RUN go mod tidy
RUN go mod vendor

# Сборка приложения
RUN go build -o myapp cmd/main.go

# Установка команды по умолчанию
CMD ["/app/myapp"]

