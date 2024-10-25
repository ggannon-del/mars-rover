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

func TestPlaceRoverOnPlanet(t *testing.T) {
	Rover := NewRover(1, 1, "N")
	planet := NewPlanet(3, 3, Rover)

	want := "ROVER"
	got := planet.GetPlanetSpot(1, 1)

	if got != want {
		t.Errorf("error placing rover on planet, got %v, expected %v", got, want)
	}

}

func TestPlaceObstacleOnGrid(t *testing.T) {
	Rover := NewRover(1, 1, "N")
	planet := NewPlanet(3, 3, Rover)
	planet.PlaceObstacle(2, 2)

	want := "O"
	got := planet.GetPlanetSpot(2, 2)

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
		Rover := NewRover(0, 0, test.startingPosition)
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
		Rover := NewRover(0, 0, test.startingPosition)
		Rover.FaceRight()
		_, _, got := Rover.GetPosition()
		if got != test.expected {
			t.Errorf("%s (%v) got %v, want %v", test.name, test.startingPosition, got, test.expected)
		}
	}
}

func TestMoveForwardHappyPath(t *testing.T) {
	tests := []struct {
		name            string
		facingDirection string
		x               int
		y               int
		expectedX       int
		expectedY       int
	}{
		{name: "Forward facing North", facingDirection: "N", x: 1, y: 1, expectedX: 1, expectedY: 2},
		{name: "Forward facing East", facingDirection: "E", x: 1, y: 1, expectedX: 2, expectedY: 1},
		{name: "Forward facing South", facingDirection: "S", x: 1, y: 1, expectedX: 1, expectedY: 0},
		{name: "Forward facing West", facingDirection: "W", x: 1, y: 1, expectedX: 0, expectedY: 1},
	}
	for _, test := range tests {
		Rover := NewRover(test.x, test.y, test.facingDirection)
		Rover.MoveForward()
		gotX, gotY, _ := Rover.GetPosition()
		if gotX != test.expectedX {
			t.Errorf("%s got %v, want %v", test.name, gotX, test.expectedX)
		}
		if gotY != test.expectedY {
			t.Errorf("%s got %v, want %v", test.name, gotY, test.expectedY)
		}
	}
}

func TestMoveForwardEdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		facingDirection string
		x               int
		y               int
		expectedX       int
		expectedY       int
	}{
		{name: "Forward facing North on upper edge of y axis", facingDirection: "N", x: 1, y: 3, expectedX: 1, expectedY: 0},
		{name: "Forward facing East on upper edge of x axis", facingDirection: "E", x: 3, y: 1, expectedX: 0, expectedY: 1},
		{name: "Forward facing South on lower edge of y axis", facingDirection: "S", x: 1, y: 0, expectedX: 1, expectedY: 3},
		{name: "Forward facing West on lower esge of x axis", facingDirection: "W", x: 0, y: 1, expectedX: 3, expectedY: 1},
	}
	for _, test := range tests {
		Rover := NewRover(test.x, test.y, test.facingDirection)
		Rover.MoveForward()
		gotX, gotY, _ := Rover.GetPosition()
		if gotX != test.expectedX {
			t.Errorf("%s got %v, want %v", test.name, gotX, test.expectedX)
		}
		if gotY != test.expectedY {
			t.Errorf("%s got %v, want %v", test.name, gotY, test.expectedY)
		}
	}
}

func TestMoveBackwardEdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		facingDirection string
		x               int
		y               int
		expectedX       int
		expectedY       int
	}{
		{name: "Backward facing North on lower edge of y axis", facingDirection: "N", x: 1, y: 0, expectedX: 1, expectedY: 3},
		{name: "Backward facing East on lower edge of x axis", facingDirection: "E", x: 0, y: 1, expectedX: 3, expectedY: 1},
		{name: "Backward facing South on upper edge of y axis", facingDirection: "S", x: 1, y: 3, expectedX: 1, expectedY: 0},
		{name: "Backward facing West on upper edge of x axis", facingDirection: "W", x: 3, y: 1, expectedX: 0, expectedY: 1},
	}
	for _, test := range tests {
		Rover := NewRover(test.x, test.y, test.facingDirection)
		Rover.MoveBackward()
		gotX, gotY, _ := Rover.GetPosition()
		if gotX != test.expectedX {
			t.Errorf("%s got %v, want %v", test.name, gotX, test.expectedX)
		}
		if gotY != test.expectedY {
			t.Errorf("%s got %v, want %v", test.name, gotY, test.expectedY)
		}
	}
}

func TestMoveBackwardHappyPath(t *testing.T) {
	tests := []struct {
		name            string
		facingDirection string
		x               int
		y               int
		expectedX       int
		expectedY       int
	}{
		{name: "Backward facing North", facingDirection: "N", x: 1, y: 1, expectedX: 1, expectedY: 0},
		{name: "Backward facing East", facingDirection: "E", x: 1, y: 1, expectedX: 0, expectedY: 1},
		{name: "Backward facing South", facingDirection: "S", x: 1, y: 1, expectedX: 1, expectedY: 2},
		{name: "Backward facing West", facingDirection: "W", x: 1, y: 1, expectedX: 2, expectedY: 1},
	}
	for _, test := range tests {
		Rover := NewRover(test.x, test.y, test.facingDirection)
		Rover.MoveBackward()
		gotX, gotY, _ := Rover.GetPosition()
		if gotX != test.expectedX {
			t.Errorf("%s got %v, want %v", test.name, gotX, test.expectedX)
		}
		if gotY != test.expectedY {
			t.Errorf("%s got %v, want %v", test.name, gotY, test.expectedY)
		}
	}
}

func TestHappyPathMoveForwardNorth(t *testing.T) {
	Rover := NewRover(1, 1, "N")

	Rover.MoveForward()

	x, y, _ := Rover.GetPosition()

	if x != 1 || y != 2 {
		t.Error("Expected position (1,2) after moving forward")
	}
}

func TestMoveBackwardNorth(t *testing.T) {
	Rover := NewRover(1, 1, "N")

	Rover.MoveBackward()

	x, y, _ := Rover.GetPosition()

	if x != 1 || y != 0 {
		t.Error("Expected position (1,2) after moving forward")
	}
}

func TestMoveForwardNorthFromTopEdgeNorthWrapsToBottomSouth(t *testing.T) {
	Rover := NewRover(1, 3, "N")

	Rover.MoveForward()
	want := 0

	_, got, _ := Rover.GetPosition()

	if got != want {
		t.Errorf("coordinate wrapping error, got %v, expected %v", got, want)
	}
}

func TestMoveBackwardNorthFromBottomEdgeSouthWrapsToTopNorth(t *testing.T) {
	Rover := NewRover(1, 0, "N")

	Rover.MoveBackward()
	want := 3

	_, got, _ := Rover.GetPosition()

	if got != want {
		t.Errorf("coordinate wrapping error, got %v, expected %v", got, want)
	}
}
