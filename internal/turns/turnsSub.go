package turns

import (
	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Turns chans
var Turns = make(map[int]map[chan *model.Turn]int)

// Add turn to chans
func Add(storyID int) chan *model.Turn {
	channel := make(chan *model.Turn, 1)
	Turns[storyID] = make(map[chan *model.Turn]int)
	Turns[storyID][channel] = 0
	return channel
}

// Update story turn
func Update(storyID int, data model.Turn) {
	for turn := range Turns[storyID] {
		turn <- &data
	}
}
