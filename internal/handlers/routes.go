package handlers

func (h *Handlers) Routes() {
	h.E.HideBanner = true

	// Добавление новой песни
	h.E.POST("/addsong", h.AddSong)
	h.E.GET("/info", h.Info)
	h.E.PUT("/update", h.Update)
	h.E.DELETE("/delete", h.Delete)

}
