package main

import "fmt"

type Point struct{ x, y int }
type Point3D struct {
	Point
	z int
}

func main() {
	var p Point3D
	p.x, p.y, p.z = 10, 20, 30
	fmt.Printf("%d\n", p.Point.x)
}
