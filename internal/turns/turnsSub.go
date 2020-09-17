package turns

import (
	"log"
	"sync"

	"github.com/scorpionknifes/pts-backend/graph/model"
)

// Turns chans
var Turns = SafeTurns{
	Turns: make(map[int]map[chan *model.Turn]int),
}

// SafeTurns safe with mutex
type SafeTurns struct {
	Turns map[int]map[chan *model.Turn]int
	Mux   sync.Mutex
}

// Add turn to chans
func Add(storyID int) chan *model.Turn {
	channel := make(chan *model.Turn, 1)
	Turns.Mux.Lock()
	if _, ok := Turns.Turns[storyID]; !ok {
		Turns.Turns[storyID] = make(map[chan *model.Turn]int)
	}

	Turns.Turns[storyID][channel] = 0
	Turns.Mux.Unlock()
	return channel
}

// Update story turn
func Update(storyID int, data model.Turn) {
	Turns.Mux.Lock()
	for turn := range Turns.Turns[storyID] {
		log.Println("send")
		turn <- &data
	}
	Turns.Mux.Unlock()
}
