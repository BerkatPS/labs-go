package main

import (
	"fmt"
	"github.com/BerkatPS/Go-Streaming/internal/usecase/genres"
	"github.com/BerkatPS/Go-Streaming/internal/usecase/movie"
	"github.com/BerkatPS/Go-Streaming/internal/usecase/shows"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
	apiKey := os.Getenv("MOVIE_APP_KEY")
	if apiKey == "" {
		log.Fatal("MOVIE_APP_KEY environtment variable is not set !!")
	}
	movieService := movie.NewService(apiKey)
	showsService := shows.NewService(apiKey)
	genreService := genres.NewService(apiKey)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Golang Fiber Web Streaming",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Testing Root Endpoint")
	})

	app.Get("/movie/:id", func(ctx *fiber.Ctx) error {
		//get Id dari url param

		movieID := ctx.Params("id")

		// convert id film to type int
		var id int
		_, err := fmt.Sscanf(movieID, "%d", &id)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Movie ID"})
		}

		getMovie, err := movieService.GetMovie(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.JSON(getMovie)
	})

	app.Get("/movies", func(ctx *fiber.Ctx) error {
		// get List Of Film
		getMovies, err := movieService.GetMovies()

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return ctx.JSON(getMovies)
	})

	app.Get("/shows", func(ctx *fiber.Ctx) error {

		getShows, err := showsService.GetShows()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return ctx.JSON(getShows)
	})

	app.Get("/genre", func(ctx *fiber.Ctx) error {

		getShows, err := genreService.GetGenres()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return ctx.JSON(getShows)
	})
	app.Get("/show/:id/episodes", func(ctx *fiber.Ctx) error {

		movieID := ctx.Params("id")

		// convert id film to type int
		var id int
		_, err := fmt.Sscanf(movieID, "%d", &id)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Shows ID"})
		}

		getShowsPerEpisode, err := showsService.GetShowsPerEpisode(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.JSON(getShowsPerEpisode)

	})
	app.Listen(":3000")

}
