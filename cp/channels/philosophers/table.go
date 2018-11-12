// Copyright 2018 Johannes Weigend
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package philosophers implements the Dining Philosophers Problem.
package philosophers

// =====================================================================================================================
// Dining Philosophers Problem. See https://en.wikipedia.org/wiki/Dining_philosophers_problem.
// The synchronization is done with channels. There are two channels for put and request fork operation.
// There is one channel per philosopher to signal when the forks can be taken.
// =====================================================================================================================

// Table represents the table with dynamic seat count.
type Table struct {
	// channel for take requests - answer is sent over the reservedCh.
	takeCh chan int
	// channel to put fork requests.
	putCh chan int
	// direct communication between philosopher and table for getting fork response.
	reservedCh []chan bool
	forkInUse  []bool
	nbrOfSeats int
}

// NewTable constructs a table with n seats.
func NewTable(nbrOfSeats int) *Table {
	table := new(Table)

	// initialize channels
	table.takeCh = make(chan int)
	table.putCh = make(chan int)

	table.reservedCh = make([]chan bool, nbrOfSeats)
	for i := 0; i < nbrOfSeats; i++ {
		table.reservedCh[i] = make(chan bool)
	}
	table.forkInUse = make([]bool, nbrOfSeats)
	table.nbrOfSeats = nbrOfSeats
	return table
}

// Function run() contains the main loop for assigning forks and starting philosophers.
func (t *Table) run() {

	for {
		select {
		case requestedFork := <-t.takeCh:
			{
				if !t.forkInUse[requestedFork] && !t.forkInUse[(requestedFork+1)%t.nbrOfSeats] {
					// both forks are not in use -> reserve
					t.forkInUse[requestedFork] = true
					t.forkInUse[(requestedFork+1)%t.nbrOfSeats] = true
					t.reservedCh[requestedFork] <- true
				} else {
					t.reservedCh[requestedFork] <- false // not valid try again --> see loop in takeForks.
				}
			}
		case putFork := <-t.putCh:
			{
				// put forks
				t.forkInUse[putFork] = false
				t.forkInUse[(putFork+1)%t.nbrOfSeats] = false
			}
		}
	}
}

func (t *Table) getTakeChannel() chan int {
	return t.takeCh
}

func (t *Table) getPutChannel() chan int {
	return t.putCh
}

func (t *Table) getReservedChannel(id int) chan bool {
	return t.reservedCh[id]
}
