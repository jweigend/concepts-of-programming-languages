// Package fan contains fan-in and fan-out functions for channel operations.
package fan

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
