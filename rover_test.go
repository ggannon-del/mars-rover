package main

import "testing"

func TestRoversInlitilisation(t *testing.T) {
	planet := NewPlanet(4, 4)

	Rover := NewRover(0, 0, "N", planet)

	x, y, _ := Rover.GetPosition()

	if x != 0 || y != 0 {
		t.Error("Expected position (0,0)", x, y)
	}
}

func TestPlaceRoverOnPlanet(t *testing.T) {
	planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", planet)

	want := "R"
	got := Rover.GetPlanetSpot(1, 1)

	if got != want {
		t.Errorf("error placing rover on planet, got %v, expected %v", got, want)
	}

}

func TestOriginalRoverPositionDoesNotContainRover(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(0, 0, "N", Planet)

	Rover.MoveForward()
	want := ""
	got := Rover.GetPlanetSpot(0, 0)

	if got != want {
		t.Errorf("error when updating grid when rover moves, got %v, expected %v", got, want)
	}
}
