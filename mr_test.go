package main

import (
	"testing"
)

func TestRoversInlitilisation(t *testing.T) {
	Rover := NewRover(0, 0, "N")

	x, y, _ := Rover.GetPosition()

	if x != 0 || y != 0 {
		t.Error("Expected position (0,0)", x, y)
	}
}

func TestFaceLeft(t *testing.T) {
	Rover := NewRover(0, 0, "N")

	Rover.FaceLeft()

	_, _, direction := Rover.GetPosition()

	if direction != "W" {
		t.Error("Expected direction W after turning left")
	}
}
