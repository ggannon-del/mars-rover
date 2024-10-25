package main

import "fmt"

type Rover struct {
	x, y      int
	direction string
}

type Planet struct {
	height, width int
	grid          [][]string
	rover         *Rover
}

const gridSize = 4

func NewRover(x, y int, direction string) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

func NewPlanet(height, width int, rover *Rover) *Planet {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}
	planet := &Planet{
		height: height,
		width:  width,
		grid:   grid,
		rover:  rover,
	}
	planet.updateRoverPositionOnGrid()
	return planet
}

func (p *Planet) PlaceObstacle(x, y int) {
	p.grid[x][y] = "O"
}

func (p *Planet) displayPlanet() {
	for _, row := range p.grid {
		for _, spot := range row {
			if spot == "" { // If the spot is empty (uninitialized), print a space or placeholder
				fmt.Print(".") // Two spaces for an empty spot for better readability
			} else {
				fmt.Print(spot, " ") // Print the current value with a space for clarity
			}
		}
		fmt.Println() // Newline after each row
	}
	fmt.Println() // Extra newline for better readability
}

func (p *Planet) updateRoverPositionOnGrid() {
	p.grid[p.rover.x][p.rover.y] = "ROVER"
}

func (p *Planet) GetPlanetSpot(x, y int) string {
	return p.grid[x][y]
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

func (r *Rover) MoveForward() {
	r.Move(1)
}

func (r *Rover) MoveBackward() {
	r.Move(-1)
}

func (r *Rover) Move(step int) {
	switch r.direction {
	case "N":
		r.y = (r.y + step + gridSize) % gridSize
	case "E":
		r.x = (r.x + step + gridSize) % gridSize
	case "S":
		r.y = (r.y - step + gridSize) % gridSize
	case "W":
		r.x = (r.x - step + gridSize) % gridSize
	}
	Planet.updateRoverPositionOnGrid()
}
