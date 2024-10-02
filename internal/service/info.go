package service

import "errors"

var (
	IDEmpty    = "не указан ID"
	ErrIDEmpty = errors.New("id empty")

	SongEmpty    = "не указана песня"
	ErrSongEmpty = errors.New("song empty")

	GroupEmpty    = "не указана группа"
	ErrGroupEmpty = errors.New("group empty")

	SongNotFound = "песня не найдена"

	GoStorage = "-база данных-"
)
