package stories

import (
	"sync"

	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Stories chans
var Stories = SafeStories{
	Stories: make(map[chan []*model.Story]int),
}

// SafeStories safe with mutex
type SafeStories struct {
	Stories map[chan []*model.Story]int
	Mux     sync.Mutex
}

// Add Story to chans
func Add() chan []*model.Story {
	Stories.Mux.Lock()
	channel := make(chan []*model.Story, 1)
	Stories.Stories[channel] = 0
	Stories.Mux.Unlock()
	return channel
}

// Update story Story
func Update(data model.Story) {
	Stories.Mux.Lock()
	for story := range Stories.Stories {
		story <- []*model.Story{&data}
	}
	Stories.Mux.Unlock()
}
