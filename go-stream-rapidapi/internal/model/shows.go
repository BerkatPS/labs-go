package model

type ShowsFilm struct {
	ID           int    `json:"_id"`
	BackdropPath string `json:"backdrop_path"`
	FirstAired   string `json:"first_aired"`
	Genre        []struct {
		Name string `json:"name"`
	} `json:"genre"`
	OriginalTitle string `json:"original_title"`
	PosterPath    string `json:"poster_path"`
	Title         string `json:"title"`
}

type ShowsFilmPerEpisode struct {
	ID            int    `json:"_id"`
	EpisodeNumber int    `json:"episode_number"`
	FirstAired    string `json:"first_aired"`
	SeasonNumber  string `json:"season_number"`
	ShowId        int    `json:"show_id"`
	ThumbnailPath string `json:"thumbnail_path"`
	Title         string `json:"title"`
	Sources       []struct {
		Netflix string `json:"netflix"`
		Hulu    string `json:"hulu"`
	} `json:"sources"`
}
