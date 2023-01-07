package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if x <= 0 {
		return nil, fmt.Errorf("x must be positive")
	}

	if y <= 0 {
		return nil, fmt.Errorf("y must be positive")
	}

	if length <= 0 {
		return nil, fmt.Errorf("length must be positive")
	}

	return &Square{
		X:      x,
		Y:      y,
		Length: length,
	}, nil
}

func (s *Square) Move(x int, y int) {
	s.X += x
	s.Y += y
}

func (s Square) Area() int {
	return s.Length * s.Length
}
