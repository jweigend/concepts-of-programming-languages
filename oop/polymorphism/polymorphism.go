package main

import (
	"fmt"
)

// Point is a two dimensional point in a cartesian coordinate system.
type Point struct{ x, y int }

// String implements the fmt.Stringer interface.
func (p Point) String() string {
	return fmt.Sprintf("x=%v,y=%v", p.x, p.y)
}

// ColorPoint extends Point by adding a color field.
type ColorPoint struct {
	Point // Embedding simulates inheritance but it is delegation!
	c     int
}

// String implements the fmt.Stringer interface.
func (p ColorPoint) String() string {
	return fmt.Sprintf("x=%v,y=%v,c=%v", p.x, p.y, p.c)
	// OR: return fmt.Sprintf("%v,c=%v", p.Point, p.c)  // Delegate to Point.String()
}

// END1 OMIT

func main() {
	var p = Point{1, 2}
	var cp = ColorPoint{Point{1, 2}, 3}

	fmt.Println(p)
	fmt.Println(cp)
	fmt.Println(cp.x) // access inherited field

	// p = cp does not work: No type hierarchy, no polymorphism

	// s is a interface and supports Polymorphism
	var s fmt.Stringer
	s = p
	fmt.Println(s.String())
	s = cp
	fmt.Println(s.String())
}

// END2 OMIT
