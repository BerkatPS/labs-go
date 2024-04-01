package shows

import (
	"encoding/json"
	"fmt"
	"github.com/BerkatPS/Go-Streaming/internal/model"
	"net/http"
)

type ShowsService struct {
	APIKey string
}

func NewService(apiKey string) *ShowsService {
	return &ShowsService{APIKey: apiKey}
}

func (ss *ShowsService) GetShows() ([]*model.ShowsFilm, error) {
	apiKey := "https://streamlinewatch-streaming-guide.p.rapidapi.com/shows?region=US&sort=popularity&sources=netflix%2Chulu&offset=0&limit=5"

	request, err := http.NewRequest("GET", apiKey, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ss.APIKey)
	request.Header.Set("X-RapidAPI-Host", "streamlinewatch-streaming-guide.p.rapidapi.com")

	client := &http.Client{}
	response, err := client.Do(request)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Api Returned non-200 status code %d", response.StatusCode)
	}

	var showsFilm []*model.ShowsFilm
	if err := json.NewDecoder(response.Body).Decode(&showsFilm); err != nil {
		return nil, err
	}

	if len(showsFilm) == 0 {
		return nil, fmt.Errorf("Shows Not found")
	}

	return showsFilm, nil
}

func (ss *ShowsService) GetShowId(showsId int) (*model.ShowsFilm, error) {
	url := fmt.Sprintf("https://streamlinewatch-streaming-guide.p.rapidapi.com/movies/%d?platform=ios&region=US", showsId)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ss.APIKey)
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

	var shows []model.ShowsFilm
	if err := json.NewDecoder(response.Body).Decode(&shows); err != nil {
		return nil, err
	}

	if len(shows) == 0 {
		return nil, fmt.Errorf("Movie not Found")
	}

	showsFilm := shows[0]

	return &showsFilm, nil
}

func (ss *ShowsService) GetShowsPerEpisode(episodeID int) (*model.ShowsFilmPerEpisode, error) {
	url := fmt.Sprintf("https://streamlinewatch-streaming-guide.p.rapidapi.com/shows/%d/episodes?platform=ios&season=1&sort=regular&offset=0&limit=25&region=US", episodeID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-RapidAPI-Key", ss.APIKey)
	request.Header.Set("X-RapidAPI-Host", "streamlinewatch-streaming-guide.p.rapidapi.com")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code %d", response.StatusCode)
	}

	var episodeData model.ShowsFilmPerEpisode
	if err := json.NewDecoder(response.Body).Decode(&episodeData); err != nil {
		return nil, err
	}

	return &episodeData, nil
}
