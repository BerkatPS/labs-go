package model

type SourcesFilm struct {
	Source      string   `json:"source"`
	DisplayName string   `json:"display_name"`
	Type        string   `json:"type"`
	Info        string   `json:"info"`
	Platform    Platform `json:"platform"`
}

type Platform struct {
	Android   string `json:"android"`
	AndroidTV string `json:"android_tv"`
	IOS       string `json:"ios"`
	Web       string `json:"web"`
}
