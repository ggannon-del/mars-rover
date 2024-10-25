package main

import (
	"testing"
)

func TestPlaceObstacleOnGrid(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", Planet)
	Rover.planet.PlaceObstacle(2, 2)

	want := "O"
	got := Rover.GetPlanetSpot(2, 2)

	if got != want {
		t.Errorf("error placing obstacle. Got %v, expected %v", got, want)
	}

}

func TestFaceLeft(t *testing.T) {
	tests := []struct {
		name             string
		startingPosition string
		expected         string
	}{
		{name: "Turn left from North", startingPosition: "N", expected: "W"},
		{name: "Turn left from East", startingPosition: "E", expected: "N"},
		{name: "Turn left from South", startingPosition: "S", expected: "E"},
		{name: "Turn left from West", startingPosition: "W", expected: "S"},
	}
	for _, test := range tests {
		Planet := NewPlanet(4, 4)

		Rover := NewRover(0, 0, test.startingPosition, Planet)
		Rover.FaceLeft()
		_, _, got := Rover.GetPosition()
		if got != test.expected {
			t.Errorf("%s (%v) got %v, want %v", test.name, test.startingPosition, got, test.expected)
		}
	}
}

func TestFaceRight(t *testing.T) {
	tests := []struct {
		name             string
		startingPosition string
		expected         string
	}{
		{name: "Turn right from North", startingPosition: "N", expected: "E"},
		{name: "Turn right from East", startingPosition: "E", expected: "S"},
		{name: "Turn right from South", startingPosition: "S", expected: "W"},
		{name: "Turn right from West", startingPosition: "W", expected: "N"},
	}
	for _, test := range tests {
		Planet := NewPlanet(4, 4)

		Rover := NewRover(0, 0, test.startingPosition, Planet)
		Rover.FaceRight()
		_, _, got := Rover.GetPosition()
		if got != test.expected {
			t.Errorf("%s (%v) got %v, want %v", test.name, test.startingPosition, got, test.expected)
		}
	}
}
