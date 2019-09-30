// Copyright 2018 by Johannes Siedersleben
// 04.07.2018

package main

import (
	"container/list"
	"fmt"
)

func main() {
	var xs = list.New()
	for i := 0; i < 10; i++ {
		xs.PushBack(i+1000)
	}

	var x list.Element
	x.Value = 7

	var y *list.Element
	y = xs.Front()
	fmt.Println(y.Value)

	var sum int
	for x := xs.Front(); x != nil; x = x.Next() {
		sum += x.Value.(int)
	}

	fmt.Print(sum)
}
