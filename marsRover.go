package main

type Rover struct {
	x, y      int
	direction string
}

const gridSize = 4

func NewRover(x, y int, direction string) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

func (r *Rover) GetPosition() (int, int, string) {
	return r.x, r.y, r.direction
}

// refactor later to array
func (r *Rover) FaceLeft() {
	if r.direction == "N" {
		r.direction = "W"
	} else if r.direction == "W" {
		r.direction = "S"
	} else if r.direction == "S" {
		r.direction = "E"
	} else if r.direction == "E" {
		r.direction = "N"
	}
}

func (r *Rover) FaceRight() {
	if r.direction == "N" {
		r.direction = "E"
	} else if r.direction == "E" {
		r.direction = "S"
	} else if r.direction == "S" {
		r.direction = "W"
	} else if r.direction == "W" {
		r.direction = "N"
	}
}

func (r *Rover) MoveForward() {
	switch r.direction {
	case "N":
		r.y = (r.y + 1) % gridSize
	case "E":
		r.x = (r.x + 1) % gridSize
	case "S":
		r.y = (r.y - 1 + gridSize) % gridSize
	case "W":
		r.x = (r.x - 1 + gridSize) % gridSize
	}
}

func (r *Rover) MoveBackward() {
	switch r.direction {
	case "N":
		r.y = (r.y - 1 + gridSize) % gridSize
	case "E":
		r.x = (r.x - 1 + gridSize) % gridSize
	case "S":
		r.y = (r.y + 1) % gridSize
	case "W":
		r.x = (r.x + 1) % gridSize
	}
}
