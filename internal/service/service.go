package service

import (
	"fmt"
	"log"
	"music/internal/models"
)

type Service struct {
	Storage StorageInt
}

type StorageInt interface {
	AddSong(song models.AddSong) (string, error)
	Info(groupName string, songTitle string) (models.SongDetail, error)
	Update(song models.Song) (string, error)
	Delete(song models.SongDel) (string, error)
	GetAllSongs(filter string) ([]models.Lib, error)
}

func New(s StorageInt) *Service {
	return &Service{
		Storage: s,
	}
}

func (s *Service) AddSong(song models.AddSong) (string, error) {
	// todo РЕГИСТР ЗАПРОСОВ
	const op = "internal.service.AddSong"

	if song.GroupName == "" {
		log.Println(op, GroupEmpty)
		return GroupEmpty, ErrGroupEmpty
	}

	if song.SongTitle == "" {
		log.Println(op, SongEmpty)
		return SongEmpty, ErrSongEmpty
	}

	log.Println(GoStorage)
	// Получаем данные из хранилища
	result, err := s.Storage.AddSong(song)
	if err != nil {
		switch err.Error() {

		case "clone":
			return result, err
		default:
			return result, err
		}
	}

	return result, nil
}

func (s *Service) Info(song models.Song) (string, error) {
	const op = "internal.service.Info"

	if song.Title == "" {
		log.Println(op, SongEmpty)
		return SongEmpty, ErrSongEmpty
	}

	if song.Group == "" {
		log.Println(op, GroupEmpty)
		return GroupEmpty, ErrGroupEmpty
	}

	log.Println(GoStorage)
	// Получаем данные из хранилища
	result, err := s.Storage.Info(song.Group, song.Title)
	if err != nil {
		log.Println("storage.Info: ", err)
		return SongNotFound, err
	}

	info := fmt.Sprintf("Дата релиза: %s, Текст: %s, Ссылка: %s", result.ReleaseDate, result.Text, result.Link)
	return info, nil
}

func (s *Service) Update(song models.Song) (string, error) {
	const op = "internal.service.Update"

	if song.Group == "" {
		log.Println(op, GroupEmpty)
		return GroupEmpty, ErrGroupEmpty
	}

	if song.Title == "" {
		log.Println(op, SongEmpty)
		return SongEmpty, ErrSongEmpty
	}

	log.Println(GoStorage)
	// Получаем данные из хранилища
	result, err := s.Storage.Update(song)
	if err != nil {
		log.Println("storage.Update: ", err)
		return result, err
	}

	return result, nil
}

func (s *Service) Delete(song models.SongDel) (string, error) {

	if song.ID == 0 {
		log.Println(IDEmpty)
		return IDEmpty, ErrIDEmpty
	}

	if song.Title == "" {
		log.Println(SongEmpty)
		return SongEmpty, ErrSongEmpty
	}

	log.Println(GoStorage)
	// Получаем данные из хранилища
	result, err := s.Storage.Delete(song)
	if err != nil {
		log.Println("storage.Delete: ", err)
		return result, err
	}

	return result, nil
}

func (s *Service) GetAllSongs(filter string) ([]models.Lib, error) {
	const op = "internal.service.GetAllSongs"

	log.Println(GoStorage)

	// if filter == "" {

	result, err := s.Storage.GetAllSongs(filter)
	if err != nil {
		return []models.Lib{}, err
	}
	return result, nil
}
