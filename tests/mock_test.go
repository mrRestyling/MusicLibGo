package tests

import (
	"log"
	"music/internal/models"
	"music/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAddSongM(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := storage.New(db)

	// отдельная структура для аргументов метода
	type args struct {
		song models.AddSong
	}

	// кастомный тип
	// функция, которая определяет объект для базы данных

	type mockBehavior func(args args, id int)

	testTable := []struct {
		name         string // имя теста
		mockBehavior mockBehavior
		args         args // входящие параметры
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				song: models.AddSong{
					GroupName:   "test",
					SongTitle:   "test",
					ReleaseDate: "test",
					Text:        "test",
					Link:        "test",
				},
			},
			id: 2,

			mockBehavior: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM groups WHERE name = $1").
					WithArgs(args.song.GroupName).
					WillReturnRows(rows)

				log.Println("1---")
				log.Println(rows)
				log.Println(err)
				log.Println()

				// Если группа не найдена, имитируем вставку новой группы
				rows = sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO groups").
					WithArgs(args.song.GroupName).
					WillReturnRows(rows)

				log.Println("2---")
				log.Println(rows)
				log.Println(err)
				log.Println()

				// Mock song does not exist
				rows = sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("SELECT id FROM songs WHERE title = $1 AND group_id = $2").
					WithArgs(args.song.SongTitle, 1).
					WillReturnRows(rows)

				log.Println("3---")
				log.Println(rows)
				log.Println(err)
				log.Println()

				// Mock insert song
				rows = sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO songs").
					WithArgs(args.song.SongTitle, 1, args.song.ReleaseDate, args.song.Text, args.song.Link).
					WillReturnRows(rows)

				log.Println("4---")
				log.Println(rows)
				log.Println(err)
				log.Println()

			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			got, err := r.AddSong(testCase.args.song)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}

		})

	}
}
