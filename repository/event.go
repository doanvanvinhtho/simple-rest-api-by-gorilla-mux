package repository

import (
	"errors"

	"github.com/doanvanvinhtho/simple-rest-api-by-gorilla-mux/model"
)

var events = []model.Event{
	{
		ID:          "6yh4tgrt4y58jdfuer",
		Title:       "Go",
		Description: "https://golang.org/",
	},
	{
		ID:          "uyhetgdtrg36223jye",
		Title:       "Mux",
		Description: "https://github.com/gorilla/mux",
	},
}

// GetOneEvent helps to get an event from repository
func GetOneEvent(id string) (*model.Event, error) {
	for _, singleEvent := range events {
		if singleEvent.ID == id {
			return &singleEvent, nil
		}
	}

	return nil, errors.New("Event " + id + " not found")
}
