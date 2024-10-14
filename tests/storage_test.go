package tests

import (
	"music/internal/models"
	"music/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStorage struct {
	service.StorageInt
}

func (s *TestStorage) AddSong(song models.AddSong) (string, error) {
	return "testSong added to testGroup", nil
}

func TestAddSong(t *testing.T) {

	// Создаем тестовую структуру
	testStorage := &TestStorage{}
	svc := service.New(testStorage)

	// Создаем тестовые данные
	testSong := models.AddSong{
		GroupName: "testGroup",
		SongTitle: "testSong",
	}

	// Вызываем метод AddSong
	result, err := svc.AddSong(testSong)

	// Проверяем, что метод вернул ожидаемые значения
	if result != "testSong added to testGroup" {
		t.Errorf("expected 'testSong added to testGroup', got %s", result)
	}
	assert.NoError(t, err)

	emptyReq := models.AddSong{}
	_, err = svc.AddSong(emptyReq)
	assert.Error(t, err)

}