package storage

import "errors"

var (
	ErrSongNotFound = errors.New("empty")
	SongNotFound    = "песня не найдена"

	ErrGroupNotFound = errors.New("empty")
	GroupNotFound    = "группа не найдена"

	ErrInternal = errors.New("internal")
	Internal    = "внутренняя ошибка"

	ErrClone = errors.New("clone")

	UpdateOK = "песня обновлена"
	NoUpdate = "песня не обновлена"

	DeleteOK = "песня удалена"
	NoDelete = "песня не удалена"

	AddGroupOK = "группа добавлена в базу данных"
	AddSongOK  = "песня добавлена в базу данных"

	AlreadySong  = "песня существует в базе данных"
	AlreadyGroup = "группа в базе данных"
)
