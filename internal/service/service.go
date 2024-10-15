package service

import (
	"fmt"
	"log"
	"music/internal/models"
	"strings"
)

type Service struct {
	Storage StorageInt
}

type StorageInt interface {
	AddSong(song models.AddSong) (string, error)
	Info(groupName string, songTitle string) (models.SongDetail, error)
	Update(song models.Song) (string, error)
	Delete(song models.SongDel) (string, error)
	GetAllSongs(filter models.Filter) (models.SongResponse, error)
	Text(info models.TextSong) (models.TextSong, error)
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

	return fmt.Sprintf("Песня добавлена в базу данных, id: %s", result), nil
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

	if song.ID == 0 {
		log.Println(op, IDEmpty)
		return IDEmpty, ErrIDEmpty
	}

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
	const op = "internal.service.Delete"

	if song.ID == 0 {
		log.Println(op, IDEmpty)
		return IDEmpty, ErrIDEmpty
	}

	if song.Title == "" {
		log.Println(op, SongEmpty)
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

func (s *Service) GetAllSongs(filter models.Filter) (models.SongResponse, error) {
	const op = "internal.service.GetAllSongs"

	log.Println(GoStorage)

	if filter.Limit == 0 {
		filter.Limit = 3
	}
	if filter.Offset == 0 {
		filter.Offset = 0
	}

	result, err := s.Storage.GetAllSongs(filter)
	if err != nil {
		log.Println(op, err)
		return models.SongResponse{}, err
	}
	return result, nil
}

func (s *Service) Text(info models.TextSong) (string, error) {
	const op = "internal.service.Text"

	log.Println(info)

	if info.Title == "" {
		log.Println(op, SongEmpty)
		return SongEmpty, ErrSongEmpty
	}

	if info.Group == "" {
		log.Println(op, GroupEmpty)
		return GroupEmpty, ErrGroupEmpty
	}

	if info.Couplet == 0 {
		info.Couplet = 1
	}

	if info.Text != "" {
		info.Text = ""
	}

	resultStorage, err := s.Storage.Text(info)
	if err != nil {
		return "не удалось вернуть текст по заданным параметрам", err
	}

	result := strings.Split(resultStorage.Text, "\n\n")

	if info.Couplet > len(result) {
		info.Couplet = len(result)
	}

	return fmt.Sprintf("Песня: %s\n Группа: %s\n Куплет: %d\n Текст: %s", info.Title, info.Group, info.Couplet, result[info.Couplet-1]), nil
}
