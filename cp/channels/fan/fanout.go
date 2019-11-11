// Package utils contains examples for fanout und fanin functions using channels.
package utils

// FanOut reads from a channel and starts an async processing task.
// The result values of the tasks will be returned in the result channel
func FanOut(input chan int, task func(int, chan int)) chan int {
	result := make(chan int)
	go func() {
		for {
			x, ok := <-input
			if !ok {
				break
			}
			go task(x, result)
		}
	}()
	return result
}

// EOF OMIT
