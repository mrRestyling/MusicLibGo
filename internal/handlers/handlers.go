package handlers

import (
	"log"
	"music/internal/models"
	"net/http"

	"github.com/labstack/echo"
)

type Handlers struct {
	E       *echo.Echo
	Service ServiceInt
}

type ServiceInt interface {
	AddSong(song models.AddSong) (string, error)
	Info(song models.Song) (string, error)
	Update(song models.Song) (string, error)
	Delete(song models.SongDel) (string, error)
}

func New(s ServiceInt) *Handlers {
	return &Handlers{
		E:       echo.New(),
		Service: s,
	}
}

func (h *Handlers) AddSong(c echo.Context) error {
	const op = "internal.handlers.AddSong"
	log.Println(op, TextAddSong)

	var song models.AddSong

	if err := c.Bind(&song); err != nil {
		log.Println(BadJSON)
		log.Printf("%s: %v", op, err)
		return c.JSON(http.StatusBadRequest, BadRequest)
	}

	log.Println(GoService)
	// Получаем данные из сервисного слоя
	result, err := h.Service.AddSong(song)
	if err != nil {
		return h.ModelError(c, err, result)
	}

	log.Println(Success)

	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) Info(c echo.Context) error {
	const op = "internal.handlers.Info"
	log.Println(op, TextInfo)

	group := c.QueryParam("group")
	song := c.QueryParam("song")

	log.Println(GoService)
	// Получаем данные из сервисного слоя
	result, err := h.Service.Info(models.Song{Group: group, Title: song})
	if err != nil {
		return h.ModelError(c, err, result)
	}

	log.Println(Success)
	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) Update(c echo.Context) error {
	const op = "internal.handlers.Update"
	log.Println(op, TextUpdate)

	var song models.Song
	if err := c.Bind(&song); err != nil {
		log.Println(BadJSON)
		log.Printf("%s: %v", op, err)
		return c.JSON(http.StatusBadRequest, BadRequest)
	}

	log.Println(GoService)
	// Получаем данные из сервисного слоя
	result, err := h.Service.Update(song)
	if err != nil {
		return h.ModelError(c, err, result)
	}

	log.Println(Success)
	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) Delete(c echo.Context) error {
	const op = "internal.handlers.Delete"
	log.Println(op, TextDelete)

	var song models.SongDel
	if err := c.Bind(&song); err != nil {
		log.Println(BadJSON)
		log.Printf("%s: %v", op, err)
		return c.JSON(http.StatusBadRequest, BadRequest)
	}

	log.Println(GoService)
	// Получаем данные из сервисного слоя
	result, err := h.Service.Delete(song)
	if err != nil {
		return h.ModelError(c, err, result)
	}

	log.Println(Success)
	return c.JSON(http.StatusOK, result)
}
