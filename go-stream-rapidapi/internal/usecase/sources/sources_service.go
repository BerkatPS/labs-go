package sources

import (
	"encoding/json"
	"fmt"
	"github.com/BerkatPS/Go-Streaming/internal/model"
	"net/http"
)

type SourcesFilm struct {
	APIKey string
}

func NewService(apiKey string) *SourcesFilm {
	return &SourcesFilm{APIKey: apiKey}
}

func (ss *SourcesFilm) GetSources() ([]*model.SourcesFilm, error) {
	apiKey := "https://streamlinewatch-streaming-guide.p.rapidapi.com/regions?region=US"

	request, err := http.NewRequest("GET", apiKey, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ss.APIKey)
	request.Header.Set("X-RapidAPI-Host", "streamlinewatch-streaming-guide.p.rapidapi.com")

	client := &http.Client{}
	response, err := client.Do(request)

	// If error. Defer to Close the response
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code %d", response.StatusCode)
	}

	var sourceFilm []*model.SourcesFilm
	if err := json.NewDecoder(response.Body).Decode(&sourceFilm); err != nil {
		return nil, err
	}

	if len(sourceFilm) == 0 {
		return nil, fmt.Errorf("Sources Film Not Found")
	}

	return sourceFilm, nil
}
