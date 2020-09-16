package stories

import (
	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Stories chans
var Stories = make([]chan *model.Story, 0)

// Add Story to chans
func Add() chan *model.Story {
	channel := make(chan *model.Story, 1)
	Stories = append(Stories, channel)
	return channel
}

// Update story Story
func Update(data model.Story) {
	for _, story := range Stories {
		story <- &data
	}
}
