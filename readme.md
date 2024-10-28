# MusicAPI.

## Технологии и инструменты

- PostgreSQL
- Docker
- Тестирование (Unit,Mock)
- Echo
- Sqlx
- Env
- JSON
- Логирование
- Graceful shutdown

## Порядок запуска приложения

1. Создайте и заполните файл `.env` (пример `.env.sample`)

2. Запустите docker-compose (Команда: `docker-compose up -d`)


## Использование пакетов:

1. пакет Эхо:
go get github.com/labstack/echo

2. пакет sqlx:
go get github.com/jmoiron/sqlx

3. драйвер Постгрес:
go get github.com/lib/pq

4. Моки для sqlx 
go get -u github.com/zhashkevych/go-sqlxmock@master



# Описание

1. Получение данных библиотеки с фильтрацией по всем полям и пагинацией
2. Получение текста песни с пагинацией по куплетам
3. Удаление песни
4. Изменение данных песни
5. Добавление новой песни в формате JSON
6. Структура БД создана путем миграций при старте сервиса
7. Логи debug- и info-
8. Конфигурационные данные в .env-файле
9. (TODO) Сваггер на реализованное АПИ 


