package storage

import (
	"log"
	"music/internal/models"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{Db: db}
}

func ConnectDB() *sqlx.DB {

	// ДЛЯ ДОКЕРА использовать эту строку:
	// db, err := sqlx.Connect("postgres", "host=postgres user="+os.Getenv("POSTGRES_USER")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable password="+os.Getenv("POSTGRES_PASSWORD"))
	db, err := sqlx.Connect("postgres", "host="+os.Getenv("HOST_SONG")+" user="+os.Getenv("POSTGRES_USER")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable password="+os.Getenv("POSTGRES_PASSWORD"))

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	return db
}

func (s *Storage) AddSong(song models.AddSong) (string, error) {
	const op = "internal.storage.AddSong"

	// Проверяем, существует ли группа в базе данных
	var groupId int64
	err := s.Db.Get(&groupId, "SELECT id FROM groups WHERE name = $1", song.GroupName)
	if err != nil {
		log.Println(op, GroupNotFound)
		err = s.Db.Get(&groupId, "INSERT INTO groups (name) VALUES ($1) RETURNING id", song.GroupName)
		if err != nil {
			log.Println(op, Internal)
			return Internal, ErrInternal
		}

		log.Println(AddGroupOK)
	}

	var existingSongId int64
	var songId int64

	// Проверяем, существует ли уже такая песня в базе данных
	err = s.Db.Get(&existingSongId, "SELECT id FROM songs WHERE title = $1 AND group_id = $2", song.SongTitle, groupId)

	if err != nil {
		log.Println(op, SongNotFound)
		err = s.Db.Get(&songId, "INSERT INTO songs (title, group_id, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			song.SongTitle, groupId, song.ReleaseDate, song.Text, song.Link)
		if err != nil {
			log.Println(op, Internal)
			return Internal, ErrInternal
		}
		log.Println(AddSongOK)
	} else {
		log.Println(AlreadySong)
		return AlreadySong, ErrClone
	}

	// return fmt.Sprintf("Песня добавлена в базу данных, id: %d", songId), nil
	return strconv.Itoa(int(songId)), nil
}

func (s *Storage) Info(groupName string, songTitle string) (models.SongDetail, error) {
	const op = "internal.storage.Info"

	// Проверяем, существует ли песня в базе данных
	var song models.SongDetail
	err := s.Db.Get(&song, "SELECT release_date, text, link FROM songs WHERE title = $1 AND group_id = (SELECT id FROM groups WHERE name = $2)",
		songTitle, groupName)
	if err != nil {
		log.Println(op, SongNotFound)
		return models.SongDetail{}, ErrSongNotFound
	}

	log.Println(AlreadySong)

	return song, nil
}

func (s *Storage) Update(song models.Song) (string, error) {
	const op = "internal.storage.Update"

	// Проверяем, существует ли песня в базе данных
	var songId int64
	err := s.Db.Get(&songId, "SELECT id FROM songs WHERE id = $1", song.ID)
	if err != nil {
		log.Println(SongNotFound)
		return SongNotFound, ErrSongNotFound
	}
	log.Println(AlreadySong)

	var groupId int64

	// Проверяем, существует ли указанная группа в базе данных
	err = s.Db.Get(&groupId, "SELECT id FROM groups WHERE name = $1", song.Group)
	if err != nil {
		log.Println(GroupNotFound)
		err = s.Db.Get(&groupId, "INSERT INTO groups (name) VALUES ($1) RETURNING id", song.Group)
		if err != nil {
			log.Println(op, Internal)
			return Internal, ErrInternal
		}
		log.Println(AlreadyGroup)
	}

	// Обновляем информацию о песне
	_, err = s.Db.Exec("UPDATE songs SET title = $1, group_id = $2, release_date = $3, text = $4, link = $5  WHERE id = $6",
		song.Title, groupId, song.ReleaseDate, song.Text, song.Link, songId)
	if err != nil {
		log.Println(op, NoUpdate)
		return NoUpdate, ErrInternal
	}
	log.Println(UpdateOK)

	return UpdateOK, nil
}

func (s *Storage) Delete(song models.SongDel) (string, error) {
	const op = "internal.storage.Delete"

	req, err := s.Db.Exec("DELETE FROM songs WHERE id = $1 AND title = $2", song.ID, song.Title)
	if err != nil {
		log.Println(op, NoDelete)
		return NoDelete, ErrInternal
	}
	rowsAffected, err := req.RowsAffected()
	if err != nil {
		log.Println(op, NoDelete)
		return NoDelete, ErrInternal
	}

	if rowsAffected == 0 {
		log.Println(SongNotFound)
		return SongNotFound, ErrSongNotFound
	}

	log.Println(DeleteOK)
	return DeleteOK, nil
}

func (s *Storage) GetAllSongs(filter models.Filter) (models.SongResponse, error) {
	const op = "internal.storage.GetAllSongs"

	log.Println("находим все песни в базе данных")

	var result []models.Lib
	var totalCount int
	var responce models.SongResponse

	query := `
    SELECT s.id, s.title, g.name AS group_name, s.release_date, s.text, s.link,
    COUNT(*) OVER () AS total_count
    FROM songs s
    JOIN groups g ON s.group_id = g.id
    WHERE ($1 = '' OR s.title ILIKE $1 OR g.name ILIKE $1 OR s.text ILIKE $1 OR s.link ILIKE $1)
    LIMIT $2 OFFSET $3
`

	rows, err := s.Db.Query(query, filter.Search, filter.Limit, filter.Offset)
	if err != nil {
		log.Println(op, "не удалось выполнить запрос")
		return responce, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Lib
		err := rows.Scan(&song.ID, &song.Title, &song.Group, &song.ReleaseDate, &song.Text, &song.Link, &totalCount)
		if err != nil {
			log.Println(op, "не удалось создать список")
			return responce, err
		}
		result = append(result, song)
	}

	responce = models.SongResponse{
		TotalCount: totalCount,
		Songs:      result,
	}

	return responce, nil
}

func (s *Storage) Text(info models.TextSong) (models.TextSong, error) {
	const op = "internal.storage.Text"

	var result models.TextSong

	query := `
    SELECT text FROM songs WHERE title = $1 AND group_id = (SELECT id FROM groups WHERE name = $2)
`
	err := s.Db.Get(&result.Text, query, info.Title, info.Group)
	if err != nil {
		log.Println(op, err)
		return models.TextSong{}, err
	}

	return result, nil
}
