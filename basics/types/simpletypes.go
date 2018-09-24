// Copyright 2018 by Johannes Siedersleben
// 29.06.2018

package main

import (
	"fmt"
	"math"
)

func nop(x int)            {} // no effect
func add(x, y int) int     { return x + y }
func addp(px, py *int) int { return *px + *py }                      // adds what px, py are pointing to
func swap(px, py *int)     { z := *px; *px = *py; *py = z }          // swaps what px, py are pointing to
func swapp(ppx, ppy **int) { var pz = *ppx; *ppx = *ppy; *ppy = pz } // swaps what ppx, ppy are pointing to

type Integer = int
type Int struct {
	value int
}

var m = 400

func id(x *int) *int { return x }
func add500()        { m += 500 }
func inc(x *int)     { *x++ }

func (x Int) Add(y Int) Int      { return NewInt(x.value + y.value) }
func NewInt(value int) Int       { var v = new(Int); v.value = value; return *v }
func Add(x, y Integer) Integer   { return x + y }
func Exec(f func())              { f() }
func Execl(f func(int), arg int) { f(arg) }

func main() {
	var c = nop
	var a = 6

	const z complex64 = complex(100, 200)
	const pi = math.Pi
	fmt.Println(real(z), imag(z), pi)

	c(a)         // calls nop
	Exec(add500) // calls add500, m = 900
	fmt.Println(m)
	Execl(func(x int) { fmt.Println(x) }, 7777) // prints 7777

	inc(&a)
	fmt.Println(a) // 7
	a = 6
	b := 8
	b = 9
	fmt.Println(a, b) // 6, 9
	swap(&a, &b)      // &a, &b unchanged, a, b swapped
	fmt.Println(a, b) // 9, 6

	var q = id(&a)
	fmt.Println(*q) // *q = 9

	var r = NewInt(78)
	var s = NewInt(43)
	var t = r.Add(s)
	fmt.Println(t) // {121}

	var u = 29
	var v = 67
	var w = Add(u, v)
	fmt.Println(w)

	pa := &a          // *pa == a
	pb := &b          // *pb == b
	swap(pa, pb)      //  a, b swapped
	fmt.Println(a, b) // 6, 9

	swap(&a, &b)          //  a, b swapped
	fmt.Println(*pa, *pb) // 9, 6

	ppa := &pa
	ppb := &pb
	swapp(ppa, ppb)       // pa, pb swapped, a, b untouched
	fmt.Println(*pa, *pb) // 6, 9
	fmt.Println(a, b)     // 9, 6

	fmt.Println(add(a, b))
	fmt.Println(addp(&a, &b))

	xs := []int{11, 22, 33}
	var sum int
	for _, x := range xs {
		sum += x
	}
	fmt.Println(sum)
}
