package main

import (
	"testing"
)

func TestRoversInlitilisation(t *testing.T) {
	Rover := NewRover(0, 0)

	x, y := Rover.GetPosition()

	if x != 0 || y != 0 {
		t.Error("Expected position (0,0)", x, y)
	}
}
