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
}
