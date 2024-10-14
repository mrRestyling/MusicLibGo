package tests

import (
	"fmt"
	"music/internal/models"
	"music/internal/service"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestService struct {
	service.StorageInt
}

func (s *TestService) AddSong(song models.AddSong) (string, error) {
	return "ADD", nil
}

func (s *TestService) Info(groupName string, songTitle string) (models.SongDetail, error) {
	return models.SongDetail{
		ReleaseDate: "TEST",
		Text:        "TEST",
		Link:        "TEST",
	}, nil

}

func TestAddSongServ(t *testing.T) {

	// Создаем тестовую структуру
	testService := &TestService{}
	svc := service.New(testService)

	// 1. Тест Happy
	testSong := models.AddSong{
		GroupName: "testGroup",
		SongTitle: "testSong",
	}

	// Вызываем метод AddSong
	result, err := svc.AddSong(testSong)

	// Проверяем, что метод вернул ожидаемые значения
	if result != "ADD" {
		t.Errorf("expected 'ADD', got %s", result)
	}
	assert.NoError(t, err)

	// 2. Тест emptyReq
	emptyReq := models.AddSong{}
	_, err = svc.AddSong(emptyReq)
	assert.Error(t, err)

	// 3. Тест emptyGroup
	emptyGroup := models.AddSong{
		SongTitle: "SongTest",
	}
	_, err = svc.AddSong(emptyGroup)
	assert.Error(t, err)

	// 4. Тест emptySong
	emptySong := models.AddSong{
		GroupName: "SongTest",
	}
	_, err = svc.AddSong(emptySong)
	assert.Error(t, err)
}

func TestInfoServ(t *testing.T) {

	testServ := &TestService{}
	srv := service.New(testServ)

	// 1. Тест Happy
	fullSong := models.Song{
		Title: "TEST",
		Group: "TEST",
	}

	answ, err := srv.Info(fullSong)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	expectedAnsw := fmt.Sprintf("Дата релиза: %s, Текст: %s, Ссылка: %s", "TEST", "TEST", "TEST")

	if !reflect.DeepEqual(answ, expectedAnsw) {
		t.Errorf("expected %v, got %v", expectedAnsw, answ)
	}
	assert.NoError(t, err)

	// 2. Тест emptyReq
	emptyReq := models.Song{}
	_, err = srv.Info(emptyReq)
	assert.Error(t, err)

	// 2+. Тест invalid song
	invalidSong := models.Song{
		Title: "",
		Group: "",
	}
	_, err = srv.Info(invalidSong)
	assert.Error(t, err)

	// 3. Тест emptyGroup
	emptyGroup := models.Song{
		Title: "TEST",
	}
	_, err = srv.Info(emptyGroup)
	assert.Error(t, err)

	// 4. Тест emptySong
	emptySong := models.Song{
		Group: "TEST",
	}
	_, err = srv.Info(emptySong)
	assert.Error(t, err)

	// 5. Тест nil
	nilSong := models.Song{}
	_, err = srv.Info(nilSong)
	assert.Error(t, err)

}
