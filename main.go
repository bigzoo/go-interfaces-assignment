package main

import "fmt"

type triangle struct {
	b float64
	h float64
}

type square struct {
	l float64
}

func (s square) getArea() float64{
	return s.l * s.l
}

func (t triangle) getArea() float64{
	return 0.5 * t.b * t.h
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{b: 4,h: 7}
	s := square{l: 4}
	printArea(t)
	printArea(s)
}
