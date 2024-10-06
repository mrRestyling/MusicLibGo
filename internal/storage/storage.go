package storage

import (
	"log"
	"music/internal/models"

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

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		// TODO: log.Println("Не удалось подключиться к базе данных")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		// TODO: log.Println("Не удалось подключиться к базе данных")
		// log.Printf("Проверьте подключение к базе данных: %s\n", connStr)
		panic(err)
	}

	return db
}

func (s *Storage) AddSong(song models.AddSong) (string, error) {
	const query = `
        WITH new_group AS (
            INSERT INTO groups (name)
            VALUES ($1)
            ON CONFLICT (name) DO UPDATE
            SET name = excluded.name
            RETURNING id
        )
        INSERT INTO songs (title, group_id, release_date, text, link)
        SELECT $2, (SELECT id FROM new_group), $3, $4, $5
        ON CONFLICT (title, group_id) DO NOTHING
    `
	_, err := s.Db.Exec(query, song.GroupName, song.SongTitle, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		return "", err
	}
	return "Song added successfully", nil
}

// // Проверяем, существует ли группа в базе данных
// var groupId int64
// err := s.Db.Get(&groupId, "SELECT id FROM groups WHERE name = $1", song.GroupName)
// if err != nil {
// 	log.Println(op, GroupNotFound)
// 	err = s.Db.Get(&groupId, "INSERT INTO groups (name) VALUES ($1) RETURNING id", song.GroupName)
// 	if err != nil {
// 		log.Println(op, Internal)
// 		return Internal, ErrInternal
// 	}

// 	log.Println(AddGroupOK)
// }

// var existingSongId int64
// var songId int64

// // Проверяем, существует ли уже такая песня в базе данных
// err = s.Db.Get(&existingSongId, "SELECT id FROM songs WHERE title = $1 AND group_id = $2", song.SongTitle, groupId)

// if err != nil {
// 	log.Println(op, SongNotFound)
// 	err = s.Db.Get(&songId, "INSERT INTO songs (title, group_id, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id",
// 		song.SongTitle, groupId, song.ReleaseDate, song.Text, song.Link)
// 	if err != nil {
// 		log.Println(op, Internal)
// 		return Internal, ErrInternal
// 	}
// 	log.Println(AddSongOK)
// } else {
// 	log.Println(AlreadySong)
// 	return AlreadySong, ErrClone
// }

// return fmt.Sprintf("Песня добавлена в базу данных, id: %d", songId), nil
// }

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
