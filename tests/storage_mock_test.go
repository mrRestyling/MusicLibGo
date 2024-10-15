package tests

import (
	"log"
	"music/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

type StorageInterface interface {
	AddSong(song models.AddSong) (string, error)
	Info(groupName string, songTitle string) (models.SongDetail, error)
	Update(song models.Song) (string, error)
	Delete(song models.SongDel) (string, error)
	GetAllSongs(filter models.Filter) (models.SongResponse, error)
	Text(info models.TextSong) (models.TextSong, error)
}

func New(s StorageInterface) StorageInterface {
	return s
}

func (m *MockStorage) AddSong(song models.AddSong) (string, error) {
	args := m.Called(song)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Info(groupName string, songTitle string) (models.SongDetail, error) {
	args := m.Called(groupName, songTitle)
	return args.Get(0).(models.SongDetail), args.Error(1)
}

func (m *MockStorage) Update(song models.Song) (string, error) {
	args := m.Called(song)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Delete(song models.SongDel) (string, error) {
	args := m.Called(song)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) GetAllSongs(filter models.Filter) (models.SongResponse, error) {
	args := m.Called(filter)
	return args.Get(0).(models.SongResponse), args.Error(1)
}

func (m *MockStorage) Text(info models.TextSong) (models.TextSong, error) {
	args := m.Called(info)
	return args.Get(0).(models.TextSong), args.Error(1)
}

func TestAddSongM(t *testing.T) {
	mockStorage := &MockStorage{}
	mockStorage.On("AddSong", models.AddSong{SongTitle: "test", GroupName: "test"}).Return("id", nil)

	storage := New(mockStorage)
	id, err := storage.AddSong(models.AddSong{SongTitle: "test", GroupName: "test"})
	log.Println("sdqdfwqf")
	assert.Equal(t, "id", id)
	assert.NoError(t, err)
}
