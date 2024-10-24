package main

type Rover struct {
	x, y int
}

func NewRover(x, y int) *Rover {
	return &Rover{
		x: x,
		y: y,
	}
}

func (r *Rover) GetPosition() (int, int) {
	return r.x, r.y
}
