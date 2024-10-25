package main

import "testing"

func TestExecuteCommandsWithoutObstacles(t *testing.T) {
	planet := NewPlanet(4, 4)
	rover := NewRover(0, 0, "N", planet)

	commands := []string{"f", "f", "r", "f"}
	err := rover.ExecuteCommands(commands)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	x, y, direction := rover.GetPosition()
	expectedX, expectedY, expectedDirection := 1, 2, "E"

	if x != expectedX || y != expectedY || direction != expectedDirection {
		t.Errorf("Expected position (%d,%d,%s), got (%d,%d,%s)", expectedX, expectedY, expectedDirection, x, y, direction)
	}
}

func TestExecuteCommandsWithObstacle(t *testing.T) {
	planet := NewPlanet(4, 4)
	planet.PlaceObstacle(1, 2)
	rover := NewRover(1, 1, "N", planet)

	commands := []string{"f", "f", "r", "f"}
	err := rover.ExecuteCommands(commands)

	if err == nil {
		t.Error("Expected an error due to obstacle, but got none")
	} else if err.Error() != "obstacle encountered at (1, 2)" {
		t.Errorf("Expected obstacle error at (1, 2), got %v", err)
	}

	x, y, _ := rover.GetPosition()
	if x != 1 || y != 1 {
		t.Errorf("Rover moved past obstacle; expected position (1,1), got (%d,%d)", x, y)
	}
}

func TestExecuteCommandsWithEdgeWrapping(t *testing.T) {
	planet := NewPlanet(4, 4)
	rover := NewRover(3, 3, "N", planet)

	commands := []string{"f", "f", "r", "f", "f"}
	err := rover.ExecuteCommands(commands)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	x, y, _ := rover.GetPosition()
	expectedX, expectedY := 1, 1

	if x != expectedX || y != expectedY {
		t.Errorf("Expected position (%d,%d) after edge wrapping, got (%d,%d)", expectedX, expectedY, x, y)
	}
}
