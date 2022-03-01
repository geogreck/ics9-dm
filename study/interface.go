package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func (p Point) Len() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

type Measurable interface {
	Len() float64
}

func main() {
	p := Point{30, 40}
	m := Measurable(p)
	fmt.Printf("%f\n", p.Len())
	fmt.Printf("%f\n", m.Len())
}
