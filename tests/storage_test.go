package tests

import (
	"log"
	"music/internal/models"
	"music/internal/storage"
	"os"
	"strconv"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddSong(t *testing.T) {

	// os.Setenv("HOST_SONG", "ADD_YOUR")         // вместо ADD_YOUR написать свои данные
	// os.Setenv("POSTGRES_USER", "ADD_YOUR")     // вместо ADD_YOUR написать свои данные
	// os.Setenv("POSTGRES_DB", "ADD_YOUR")       // вместо ADD_YOUR написать свои данные
	// os.Setenv("POSTGRES_PASSWORD", "ADD_YOUR") // вместо ADD_YOUR написать свои данные

	db, err := sqlx.Connect("postgres", "host="+os.Getenv("HOST_SONG")+" user="+os.Getenv("POSTGRES_USER")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable password="+os.Getenv("POSTGRES_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	s := storage.New(db)

	song := models.AddSong{
		SongTitle:   "Test Song",
		GroupName:   "1Test Group",
		ReleaseDate: "2022-01-01",
		Text:        "Test text",
		Link:        "https://example.com",
	}
	idStr, err := s.AddSong(song)
	if err != nil {
		t.Errorf("ошибка добавления песни: %v", err)
	}

	idInt, _ := strconv.Atoi(idStr)

	delAdd := models.SongDel{
		ID:    idInt,
		Title: song.SongTitle,
	}

	_, err = s.Delete(delAdd)
	if err != nil {
		t.Log("Не удалось удалить песню, после тестового добавления")
		require.Error(t, err)
	}
	assert.NoError(t, err)

}
