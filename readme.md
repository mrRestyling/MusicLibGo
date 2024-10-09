# MusicAPI


1. Создайте и заполните файл `.env` (пример `.env.sample`)

2. Запустите docker-compose (Команда: `docker-compose up -d`)


Использование пакетов:

1. добавляем пакет Эхо:
go get github.com/labstack/echo

2. добавляем пакет БД:
go get github.com/jmoiron/sqlx

3. добавляем драйвер Постгрес:
go get github.com/lib/pq



# musicLibGo

# ЗАДАЧА

1. Выставить rest методы
Получение данных библиотеки с фильтрацией по всем полям и пагинацией
Получение текста песни с пагинацией по куплетам
Удаление песни
Изменение данных песни
Добавление новой песни в формате:

JSON

{
 "group": "Muse",
 "song": "Supermassive Black Hole"
}

2. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)
3. Покрыть код debug- и info-логами
4. Вынести конфигурационные данные в .env-файл
5. Сгенерировать сваггер на реализованное АПИ

6. При добавлении сделать запрос в АПИ, описанного сваггером

openapi: 3.0.3
info:
  title: Music info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: group
          in: query
          required: true
          schema:
            type: string
        - name: song
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SongDetail'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    SongDetail:
      required:
        - releaseDate
        - text
        - link
      type: object
      properties:
        releaseDate:
          type: string
          example: 16.07.2006
        text:
          type: string
          example: Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight
        link:
          type: string
          example: https://www.youtube.com/watch?v=Xsp3_a-PMTw


