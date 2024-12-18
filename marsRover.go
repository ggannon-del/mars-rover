package main

import (
	"fmt"
)

type Rover struct {
	x, y      int
	direction string
	planet    *Planet
}

type Planet struct {
	height, width int
	grid          [][]string
}

func NewRover(x, y int, direction string, planet *Planet) *Rover {
	rover := &Rover{
		x:         x,
		y:         y,
		direction: direction,
		planet:    planet,
	}
	rover.updateRoverPositionOnGrid()
	return rover
}

func NewPlanet(height, width int) *Planet {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}
	return &Planet{
		height: height,
		width:  width,
		grid:   grid,
	}

}

func (p *Planet) PlaceObstacle(x, y int) {
	p.grid[x][y] = "O"
}

func (r *Rover) updateRoverPositionOnGrid() {
	r.planet.grid[r.x][r.y] = "R"
}

func (r *Rover) GetPlanetSpot(x, y int) string {
	return r.planet.grid[x][y]
}

func (r *Rover) GetPosition() (int, int, string) {
	return r.x, r.y, r.direction
}

func (r *Rover) FaceLeft() {
	directionMap := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	r.direction = directionMap[r.direction]
}

func (r *Rover) FaceRight() {
	directionMap := map[string]string{
		"N": "E",
		"W": "N",
		"S": "W",
		"E": "S",
	}
	r.direction = directionMap[r.direction]
}

func (r *Rover) MoveForward() error {
	return r.Move(1)
}

func (r *Rover) MoveBackward() error {
	return r.Move(-1)
}

func (r *Rover) clearOldPosition(x, y int) {
	r.planet.grid[x][y] = ""
}

func (r *Rover) Move(step int) error {
	r.clearOldPosition(r.x, r.y)

	newX, newY := r.calculateNewPosition(step)

	if r.isObstacleAt(newX, newY) {

		return fmt.Errorf("obstacle encountered at (%d, %d)", newX, newY)
	}

	r.x, r.y = newX, newY
	r.updateRoverPositionOnGrid()
	return nil
}
func (r *Rover) calculateNewPosition(step int) (int, int) {
	newX, newY := r.x, r.y

	switch r.direction {
	case "N":
		newY = (r.y + step) % r.planet.height
	case "E":
		newX = (r.x + step) % r.planet.width
	case "S":
		newY = (r.y - step) % r.planet.height
	case "W":
		newX = (r.x - step) % r.planet.width
	}

	if newX < 0 {
		newX += r.planet.width
	}
	if newY < 0 {
		newY += r.planet.height
	}

	return newX, newY
}

func (r *Rover) isObstacleAt(newX, newY int) bool {
	return r.planet.grid[newX][newY] == "O"
}

func (r *Rover) ExecuteCommands(commands []string) error {
	for _, command := range commands {
		var err error
		switch command {
		case "f":
			err = r.MoveForward()
		case "b":
			err = r.MoveBackward()
		case "l":
			r.FaceLeft()
		case "r":
			r.FaceRight()
		}

		if err != nil {
			return err
		}
	}
	return nil
}
