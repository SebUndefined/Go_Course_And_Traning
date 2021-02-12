package main

import (
	"fmt"
)

// Shape interface
type Shape interface {
	getArea() (float64, error)
}

// Square struct
type Square struct {
	sideLength float64
}

// Triangle struct
type Triangle struct {
	height float64
	base   float64
}

func main() {
	tr := Triangle{12.43, 3.00}
	tr2 := Triangle{0, 3.00}
	printArea(tr)
	printArea(tr2)
}

func (s Square) getArea() (float64, error) {
	if s.sideLength <= 0 {
		return 0, fmt.Errorf("Sidelength has to be upper than 0, value got %v", s.sideLength)
	}
	res := s.sideLength * s.sideLength
	return res, nil
}

func (t Triangle) getArea() (float64, error) {
	if t.base <= 0 {
		return 0, fmt.Errorf("Base has to be upper than 0, value got %v", t.base)
	}
	if t.height <= 0 {
		return 0, fmt.Errorf("Height has to be upper than 0, value got %v", t.base)
	}
	res := 0.5 * t.base * t.height
	return res, nil
}

func printArea(s Shape) {
	ar, err := s.getArea()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The area of the shape is ", ar)
	}

}
