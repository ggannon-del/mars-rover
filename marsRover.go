package main

type Rover struct {
	x, y      int
	direction string
}

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
