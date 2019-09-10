# Exercise 6 - Concurrent Programming in Go

If you do not finish during the lecture period, please finish it as homework.

## Exercise 6.1 - Warm Up

Test the Ping-Pong Programm (see slides)

## Exercise 6.2 - Fan-Out / Fan-In

Write tests for the FanOut and FanIn functions (see slides)

## Exercise 6.3 - Dining Philosophers

Write a program to simulate the Dining Philosophers Problem by using Go Channels.
- There should be one Go Routine for each Philosopher and the Table
- Make sure that:
  - The distribution of forks is fair - No philosopher dies on starvation 
  - Use the given Unit Test:

```go
func TestPhilosophers(t *testing.T) {

	var COUNT = 5

	// start table for 5 philosophers
	table := NewTable(COUNT)

	// create 5 philosophers and run parallel 
	for i := 0; i < COUNT; i++ {
		philosopher := NewPhilosopher(i, table)
		go philosopher.run()
	}
	go table.run()

	// simulate 10 milliseconds --> check output
	time.Sleep(10 * time.Millisecond)
}
```

Sample console output:

```sh
[->]: Philosopher #0 eats ...
[->]: Philosopher #3 eats ...
[<-]: Philosopher #0  eat ends.
[<-]: Philosopher #3  eat ends.
[->]: Philosopher #0 thinks ...
[->]: Philosopher #1 eats ...
[->]: Philosopher #3 thinks ...
[->]: Philosopher #4 eats ...
[<-]: Philosopher #1  eat ends.
[->]: Philosopher #1 thinks ...
[<-]: Philosopher #4  eat ends.
[->]: Philosopher #2 eats ...
[->]: Philosopher #4 thinks ...
[<-]: Philosopher #0 thinking ends
[->]: Philosopher #0 eats ...
[<-]: Philosopher #3 thinking ends
```