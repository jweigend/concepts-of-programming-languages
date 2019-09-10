package main

import (
	"fmt"
)

// Position is the position of the screen.
type Position struct {
	x, y int
}

// GeoObject is the "BaseClass" of any geometrical object. All objects have a position and a color.
type GeoObject struct {
	pos   Position
	color int
}

// Circle is a concrete GeoObject with a radius.
type Circle struct {
	GeoObject
	radius int
}

// Rectangle is a concrete GeoObject with a width and a height.
type Rectangle struct {
	GeoObject
	width, height int
}

// Triangle is a concrete GeoObject with three points (ABC).
// The coordinates of the three points are relative to the position of the object.
type Triangle struct {
	GeoObject
	p1, p2, p3 Position
}

// Painter is used to paint GeoObjects.
type Painter interface {
	Paint()
}

// Paint is implemented by Circle
func (c Circle) Paint() {
	fmt.Printf("Painting circle with radius=%v at position=%v and color=%v\n", c.radius, c.pos, c.color)
}

// Paint is implemented by Rectangle
func (r Rectangle) Paint() {
	fmt.Printf("Painting rectangle with width=%v, height=%v at position=%v and color=%v\n", r.width, r.height, r.pos, r.color)
}

// Paint is implemented by Triangle
func (c Triangle) Paint() {
	fmt.Printf("Painting triangle with p1=%v, p2=%v, p3=%v at position=%v and color=%v\n", c.p1, c.p2, c.p3, c.pos, c.color)
}

func main() {
	// Polymorph slice of Painter objects
	objects := []Painter{
		// short initialization
		Circle{GeoObject{Position{1, 2}, 3}, 40},
		Rectangle{GeoObject{Position{1, 2}, 4}, 10, 10},
		Triangle{GeoObject{Position{1, 2}, 3}, Position{10, 20}, Position{11, 21}, Position{12, 22}},

		// or with named identifiers
		Circle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 3},
			radius: 40},
		Rectangle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 4},
			width:  10,
			height: 10},
		Triangle{
			GeoObject: GeoObject{
				pos:   Position{x: 1, y: 2},
				color: 3},
			p1: Position{x: 10, y: 20},
			p2: Position{x: 11, y: 21},
			p3: Position{x: 12, y: 22}},
	}
	for _, v := range objects {
		v.Paint()
	}
}
