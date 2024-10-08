package handlers

func (h *Handlers) Routes() {
	h.E.HideBanner = true

	// Информация о песне
	h.E.GET("/info", h.Info)

	// Получение данных библиотеки с фильтрацией по всем полям и пагинацией
	h.E.GET("/infoall", h.InfoAll)

	// Получение текста песни с пагинацией по куплетам
	h.E.GET("/text", h.Text)

	// Добавление новой песни
	h.E.POST("/addsong", h.AddSong)

	// Обновление песни
	h.E.PUT("/update", h.Update)

	// Удаление песни
	h.E.DELETE("/delete", h.Delete)

}
