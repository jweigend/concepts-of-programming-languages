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
	"testing"
	"time"
)

// =====================================================================================================================
// Main Test.
// =====================================================================================================================

func TestPhilosophers(t *testing.T) {

	var COUNT = 5

	// start table for 5
	table := NewTable(COUNT)

	// start philosophers
	for i := 0; i < COUNT; i++ {
		philosopher := NewPhilosopher(i, table)
		go philosopher.run()
	}

	// wait 1 millisecond --> check output
	time.Sleep(60000 * time.Millisecond)
}
