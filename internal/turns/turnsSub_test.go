package turns

import (
	"testing"

	"github.com/scorpionknifes/pts-backend/graph/model"
)

func init() {
	Turns = make(map[int]map[chan *model.Turn]int)
}

func TestAdd(t *testing.T) {
	channel := Add(1)
	if _, ok := Turns[1][channel]; !ok {
		t.Errorf("Add() failed to add")
	}
}

func TestUpdate(t *testing.T) {
	channel := Add(1)
	Update(1, model.Turn{Value: "test"})
	test := <-channel
	if test.Value != "test" {
		t.Errorf("Update() has failed update correct data")
	}
}
