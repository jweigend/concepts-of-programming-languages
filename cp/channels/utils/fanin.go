// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package utils contains functions to work with channels.
package utils

import (
	"fmt"
)

// FanIn reads from N-Channels and forwards the result to the output channel.
func FanIn(channels []chan int, output chan int) {
	for i := 0; i < len(channels); i++ {
		// fan in
		go func(i int) {
			for {
				n, ok := <-channels[i]
				if !ok {
					break
				}
				output <- n
			}
			fmt.Println("input channel closed: done.")
		}(i)
	}
}

// AsyncReadAndPrintFromCh prints all data from a channel until the channel is closed.
func AsyncReadAndPrintFromCh(ch chan int) {
	go func() {
		for {
			res, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(res)
		}
		fmt.Println("output channel closed: done.")
	}()
}
