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

package philosophers

import (
	"fmt"
	"time"
)

// Philosopher represents one philosopher.
type Philosopher struct {
	id    int
	table *Table
}

// NewPhilosopher constructs a philosopher.
func NewPhilosopher(id int, table *Table) *Philosopher {
	p := new(Philosopher)
	p.id = id
	p.table = table
	return p
}

// Run loops forever.
func (p *Philosopher) run() {
	for {
		p.takeForks()
		p.eat()
		p.putForks()
		p.think()
	}
}

// Take forks by channeling our id to the table and wait until the table returns true on the reserved channel.
func (p *Philosopher) takeForks() {
	// try to get forks from table
	gotForks := false
	for !gotForks {
		p.table.getTakeChannel() <- p.id
		gotForks = <-p.table.getReservedChannel(p.id)
	}
}

// Put forks by channeling our id to the table. The table is responsible for the put logic.
func (p *Philosopher) putForks() {
	p.table.getPutChannel() <- p.id
}

// Eating.
func (p *Philosopher) eat() {
	fmt.Printf("[->]: Philosopher #%d eats ...\n", p.id)
	time.Sleep(2 * time.Millisecond)
	fmt.Printf("[<-]: Philosopher #%d  eat ends.\n", p.id)
}

// Thinking.
func (p *Philosopher) think() {
	fmt.Printf("[->]: Philosopher #%d thinks ...\n", p.id)
	time.Sleep(3 * time.Millisecond)
	fmt.Printf("[<-]: Philosopher #%d thinking ends\n", p.id)
}
