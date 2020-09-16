package turns

import (
	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Turns chans
var Turns = make(map[int][]chan *model.Turn)

// Add turn to chans
func Add(storyID int) chan *model.Turn {
	channel := make(chan *model.Turn, 1)
	Turns[storyID] = append(Turns[storyID], channel)
	return channel
}

// Update story turn
func Update(storyID int, data model.Turn) {
	for _, turn := range Turns[storyID] {
		turn <- &data
	}
}
