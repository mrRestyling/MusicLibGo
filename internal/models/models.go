package models

type Group struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	ID          uint   `json:"id"`
	Title       string `json:"song"`
	Group       string `json:"group"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddSong struct {
	GroupName   string `json:"group"`
	SongTitle   string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongDetail struct {
	ReleaseDate string `db:"release_date" json:"release_date"` // ??? TODO
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongDel struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
