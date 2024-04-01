package genres

import (
	"encoding/json"
	"fmt"
	"github.com/BerkatPS/Go-Streaming/internal/model"
	"net/http"
)

type GenreService struct {
	APIKey string
}

func NewService(apikey string) *GenreService {
	return &GenreService{APIKey: apikey}
}

func (gs *GenreService) GetGenres() ([]*model.GenresFilm, error) {
	apiUrl := "https://streamlinewatch-streaming-guide.p.rapidapi.com/movies?region=US&sort=popularity&sources=netflix%2Chulu&offset=0&limit=5"

	request, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", gs.APIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Returned non-200 status code %d", response.StatusCode)
	}

	var genres []*model.GenresFilm
	if err := json.NewDecoder(response.Body).Decode(&genres); err != nil {
		return nil, err
	}

	return genres, nil
}
