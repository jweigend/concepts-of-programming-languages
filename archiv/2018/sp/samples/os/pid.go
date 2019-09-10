package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println("process id: ", pid)
}
