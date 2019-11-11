// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package utils contains functions to work with channels.
package utils

import (
	"fmt"
	"log"
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
			log.Println("input channel closed: done.")
		}(i)
	}
}

// FanIn OMIT

// Print prints all data from a channel until the channel is closed.
func Print(ch chan int) {
	go func() {
		for {
			res, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(res)
		}
		log.Println("output channel closed: done.")
	}()
}
