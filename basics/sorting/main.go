package main

import (
	"log"
	"sort"
)

// AxisSorter sorts planets by axis.
type AxisSorter []Planet

func (a AxisSorter) Len() int           { return len(a) }
func (a AxisSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AxisSorter) Less(i, j int) bool { return a[i].Axis < a[j].Axis }

// NameSorter sorts planets by name.
type NameSorter []Planet

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Planet represents a planet in our solarsystem.
type Planet struct {
	Name       string  `json:"name"`
	Aphelion   float64 `json:"aphelion"`   // in million km
	Perihelion float64 `json:"perihelion"` // in million km
	Axis       int64   `json:"axis"`       // in km
	Radius     float64 `json:"radius"`
}

func main() {
	var mars Planet
	mars.Name = "Mars"
	mars.Aphelion = 249.2
	mars.Perihelion = 206.7
	mars.Axis = 227939100
	mars.Radius = 3389.5

	var earth Planet
	earth.Name = "Earth"
	earth.Aphelion = 151.930
	earth.Perihelion = 147.095
	earth.Axis = 149598261
	earth.Radius = 6371.0

	var venus Planet
	venus.Name = "Venus"
	venus.Aphelion = 108.939
	venus.Perihelion = 107.477
	venus.Axis = 108208000
	venus.Radius = 6051.8

	planets := []Planet{mars, venus, earth}
	log.Println("unsorted:", planets)

	sort.Sort(AxisSorter(planets))
	log.Println("by axis:", planets)

	sort.Sort(NameSorter(planets))
	log.Println("by name:", planets)
}
