package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLenght float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLenght: 10}
	t := triangle{height: 3, base: 5}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}
