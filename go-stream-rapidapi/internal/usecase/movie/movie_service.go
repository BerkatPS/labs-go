package movie

import (
	"encoding/json"
	"fmt"
	"github.com/BerkatPS/Go-Streaming/internal/model"
	"net/http"
)

type MovieService struct {
	APIKey string
}

func NewService(apikey string) *MovieService {
	return &MovieService{APIKey: apikey}
}

func (ms *MovieService) GetMovie(movieID int) (*model.Movie, error) {
	url := fmt.Sprintf("https://streamlinewatch-streaming-guide.p.rapidapi.com/movies/%d?platform=ios&region=US", movieID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ms.APIKey)
	request.Header.Set("X-RapidAPI-Host", "streamlinewatch-streaming-guide.p.rapidapi.com")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Returned non-200 status code %d", response.StatusCode)
	}

	var movies []model.Movie
	if err := json.NewDecoder(response.Body).Decode(&movies); err != nil {
		return nil, err
	}

	if len(movies) == 0 {
		return nil, fmt.Errorf("Movie not Found")
	}

	movie := movies[0]

	return &movie, nil
}

func (ms *MovieService) GetMovies() ([]*model.Movie, error) {
	apiUrl := "https://streamlinewatch-streaming-guide.p.rapidapi.com/movies?region=US&sort=popularity&sources=netflix%2Chulu&offset=0&limit=5"

	request, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ms.APIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Returned non-200 status code %d", response.StatusCode)
	}

	var movie []*model.Movie
	if err := json.NewDecoder(response.Body).Decode(&movie); err != nil {
		return nil, err
	}

	return movie, nil
}
