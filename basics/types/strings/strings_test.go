package strings

import (
	"fmt"
	"testing"
)

func TestStringLength(t *testing.T) {

	// Unicode code length
	length := len("A")
	fmt.Printf("Length: %d\n", length)
	length = len("±")
	fmt.Printf("Length: %d\n", length)
	length = len("☯")
	fmt.Printf("Length: %d\n", length)
}
