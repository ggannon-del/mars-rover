package main

import "testing"

func TestMovingForwardUpdatesPlanet(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", Planet)

	Rover.MoveForward()
	want := "R"
	got := Rover.GetPlanetSpot(1, 2)

	if got != want {
		t.Errorf("error when updating grid when rover moves, got %v, expected %v", got, want)
	}

}

func TestMovingBackwardUpdatesPlanet(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", Planet)

	Rover.MoveBackward()
	want := "R"
	got := Rover.GetPlanetSpot(1, 0)

	if got != want {
		t.Errorf("error when updating grid when rover moves, got %v, expected %v", got, want)
	}

}

func TestMovingForwardUpdatesPlanetEdgeCase(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(0, 3, "N", Planet)

	Rover.MoveForward()
	want := "R"
	got := Rover.GetPlanetSpot(0, 0)

	if got != want {
		t.Errorf("error when updating grid when rover moves, got %v, expected %v", got, want)
	}

}

func TestMovingBackwardUpdatesPlanetEdgeCase(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(0, 0, "N", Planet)

	Rover.MoveBackward()
	want := "R"
	got := Rover.GetPlanetSpot(0, 3)

	if got != want {
		t.Errorf("error when updating grid when rover moves, got %v, expected %v", got, want)
	}

}
