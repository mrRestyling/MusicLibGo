package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	TextAddSong = "запрос на добавление песни"
	TextInfo    = "запрос на получение информации о песне"
	TextUpdate  = "запрос на обновление песни"
	TextDelete  = "запрос на удаление песни"

	BadRequest = "пришел запрос в неверном формате"

	BadJSON = "не удалось распарсить JSON"

	GoService = "-сервисный уровень-"

	Success = "успешно отправлен ответ"
)

func (h *Handlers) ModelError(c echo.Context, err error, result string) error {

	switch err.Error() {
	case "clone":
		return c.JSON(http.StatusConflict, result)
	case "empty":
		return c.JSON(http.StatusNotFound, result)
	case "internal":
		return c.JSON(http.StatusInternalServerError, result)
	case "id empty":
		return c.JSON(http.StatusBadRequest, result)
	case "group empty":
		return c.JSON(http.StatusBadRequest, result)
	case "song empty":
		return c.JSON(http.StatusBadRequest, result)
	case "bad request":
		return c.JSON(http.StatusBadRequest, result)
	default:
		return c.JSON(http.StatusBadRequest, result)
	}

}
