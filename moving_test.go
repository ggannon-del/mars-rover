package main

import "testing"

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
		Planet := NewPlanet(4, 4)

		Rover := NewRover(test.x, test.y, test.facingDirection, Planet)

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
		Planet := NewPlanet(4, 4)

		Rover := NewRover(test.x, test.y, test.facingDirection, Planet)

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
		Planet := NewPlanet(4, 4)

		Rover := NewRover(test.x, test.y, test.facingDirection, Planet)
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
		Planet := NewPlanet(4, 4)

		Rover := NewRover(test.x, test.y, test.facingDirection, Planet)
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
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", Planet)

	Rover.MoveForward()

	x, y, _ := Rover.GetPosition()

	if x != 1 || y != 2 {
		t.Error("Expected position (1,2) after moving forward")
	}
}

func TestMoveBackwardNorth(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 1, "N", Planet)

	Rover.MoveBackward()

	x, y, _ := Rover.GetPosition()

	if x != 1 || y != 0 {
		t.Error("Expected position (1,2) after moving forward")
	}
}

func TestMoveForwardNorthFromTopEdgeNorthWrapsToBottomSouth(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 3, "N", Planet)

	Rover.MoveForward()
	want := 0

	_, got, _ := Rover.GetPosition()

	if got != want {
		t.Errorf("coordinate wrapping error, got %v, expected %v", got, want)
	}
}

func TestMoveBackwardNorthFromBottomEdgeSouthWrapsToTopNorth(t *testing.T) {
	Planet := NewPlanet(4, 4)

	Rover := NewRover(1, 0, "N", Planet)

	Rover.MoveBackward()
	want := 3

	_, got, _ := Rover.GetPosition()

	if got != want {
		t.Errorf("coordinate wrapping error, got %v, expected %v", got, want)
	}
}
