package stories

import (
	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Stories chans
var Stories = make(map[chan *model.Story]int)

// Add Story to chans
func Add() chan *model.Story {
	channel := make(chan *model.Story, 1)
	Stories[channel] = 0
	return channel
}

// Update story Story
func Update(data model.Story) {
	for story := range Stories {
		story <- &data
	}
}
