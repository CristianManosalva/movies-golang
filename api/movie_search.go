package api

import (
	"time"
)

type MovieFilter struct {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}

type Movie struct {
	Title       string    `json:"title"`
	Cast        string    `json:"cast"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       string    `json:"genre"`
	Director    string    `json:"director"`
}

type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}

type MovieService struct {
}

func (s MovieService) Search(filter MovieFilter) ([]Movie, error) {
	m1 := Movie{
		Title:       "Blande Runner",
		Cast:        "Harrison Ford",
		ReleaseDate: time.Now(),
		Genre:       "Cs Fictio",
		Director:    "Marato Devino",
	}

	m2 := Movie{
		Title:       "Blande Runner",
		Cast:        "Harrison Ford",
		ReleaseDate: time.Now(),
		Genre:       "Cs Fictio",
		Director:    "Marato Devino",
	}

	m3 := Movie{
		Title:       "Blande Runner",
		Cast:        "Harrison Ford",
		ReleaseDate: time.Now(),
		Genre:       "Cs Fictio",
		Director:    "Marato Devino",
	}

	var _movies []Movie

	_movies = append(_movies, m1, m2, m3)

	return _movies, nil
}
