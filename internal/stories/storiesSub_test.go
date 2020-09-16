package stories

import (
	"testing"

	"github.com/scorpionknifes/pts-backend/graph/model"
)

func init() {
	Stories = make(map[chan *model.Story]int)
}

func TestAdd(t *testing.T) {
	channel := Add()
	if _, ok := Stories[channel]; !ok {
		t.Errorf("Add() failed to add")
	}
}

func TestUpdate(t *testing.T) {
	channel := Add()
	Update(model.Story{Name: "test"})
	test := <-channel
	if test.Name != "test" {
		t.Errorf("Update() has failed update correct data")
	}
}