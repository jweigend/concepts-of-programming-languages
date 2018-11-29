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

import (
	"github.com/jweigend/concepts-of-programming-languages/cp/locks/resourcemanager"
)

// =====================================================================================================================
// Dining Philosophers Problem. See https://en.wikipedia.org/wiki/Dining_philosophers_problem.
// The synchronization is done with channels. There are two channels for put and request fork operation.
// There is one channel per philosopher to signal when the forks can be taken.
// =====================================================================================================================

// Table represents the table with dynamic seat count.
type Table struct {
	manager    *resourcemanager.ResourceManager
	forkInUse  []bool
	nbrOfSeats int
}

// NewTable constructs a table with n seats.
func NewTable(nbrOfSeats int) *Table {
	table := new(Table)
	table.manager = resourcemanager.NewResourceManager()
	table.forkInUse = make([]bool, nbrOfSeats)
	table.nbrOfSeats = nbrOfSeats
	return table
}

// GetManager returns the resource manager.
func (t *Table) GetManager() *resourcemanager.ResourceManager {
	return t.manager
}
