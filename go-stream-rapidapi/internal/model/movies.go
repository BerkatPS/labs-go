package model

type Movie struct {
	ID            int64    `json:"_id"`
	BackdropPath  string   `json:"backdrop_path"`
	Genres        []string `json:"genres"`
	OriginalTitle string   `json:"original_title"`
	Overview      string   `json:"overview"`
	PosterPath    string   `json:"poster_path"`
	ReleaseDate   string   `json:"release_date"`
	Title         string   `json:"title"`
}
