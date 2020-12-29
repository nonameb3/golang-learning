package main

import "fmt"

type shape interface {
	getArea() float64
}

type trianle struct {
	base   float64
	heigth float64
}
type square struct {
	lenght float64
	heigth float64
}

func main() {
	myTrianle := trianle{5, 10}
	mySquare := square{10, 10}

	shape := []shape{}
	shape = append(shape, myTrianle)
	shape = append(shape, mySquare)

	for _, data := range shape {
		printAera(data)
	}
}

func printAera(s shape) {
	fmt.Printf("%f", s.getArea())
}

func (t trianle) getArea() float64 {
	aera := t.base * t.heigth * 0.5
	return aera
}

func (s square) getArea() float64 {
	aera := s.lenght * s.heigth
	return aera
}
