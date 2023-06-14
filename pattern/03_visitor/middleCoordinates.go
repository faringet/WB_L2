package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(s *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
